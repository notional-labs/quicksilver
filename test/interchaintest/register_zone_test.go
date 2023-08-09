package interchaintest

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/icza/dyno"
	istypes "github.com/ingenuity-build/quicksilver/x/interchainstaking/types"
	"github.com/strangelove-ventures/interchaintest/v7"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
	"github.com/strangelove-ventures/interchaintest/v7/relayer"
	"github.com/strangelove-ventures/interchaintest/v7/testreporter"
	"github.com/strangelove-ventures/interchaintest/v7/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
	// simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
)

const (
	heightDelta      = 20
	votingPeriod     = "30s"
	maxDepositPeriod = "10s"
)

// Spin up a quicksilverd chain, push a contract, and get that contract code from chain. Submit a proposal to register zones and query zones.
func TestRegisterZone(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	client, network := interchaintest.DockerSetup(t)

	rep := testreporter.NewNopReporter()
	eRep := rep.RelayerExecReporter(t)

	ctx := context.Background()

	// Get both chains
	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			ChainConfig: ibc.ChainConfig{
				Type:           "cosmos",
				Name:           "quicksilver",
				ChainID:        "quicksilverd",
				Images:         []ibc.DockerImage{QuicksilverImage},
				Bin:            "quicksilverd",
				Bech32Prefix:   "quick",
				Denom:          "stake",
				GasPrices:      "0.00stake",
				GasAdjustment:  1.3,
				TrustingPeriod: "504h",
				EncodingConfig: quicksilverEncoding(),
				NoHostMount:    true,
				ModifyGenesis:  modifyGenesisShortProposals(votingPeriod, maxDepositPeriod),
			},
		},
		{
			Name:    "juno",
			Version: "v14.1.0",
		},
	})

	t.Logf("Calling cf.Chains")
	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)

	quicksilverd, juno := chains[0].(*cosmos.CosmosChain), chains[1].(*cosmos.CosmosChain)

	// Get a relayer instance
	r := interchaintest.NewBuiltinRelayerFactory(ibc.CosmosRly, zaptest.NewLogger(t), relayer.CustomDockerImage("ghcr.io/notional-labs/cosmos-relayer", "nguyen-v2.3.1", "1000:1000")).Build(t, client, network)

	// Build the network; spin up the chains and configure the relayer
	t.Logf("NewInterchain")
	ic := interchaintest.NewInterchain().
		AddChain(quicksilverd).
		AddChain(juno).
		AddRelayer(r, "rly").
		AddLink(interchaintest.InterchainLink{
			Chain1:  quicksilverd,
			Chain2:  juno,
			Relayer: r,
			Path:    pathQuicksilverJuno,
		})

	t.Logf("Interchain build options")
	require.NoError(t, ic.Build(ctx, eRep, interchaintest.InterchainBuildOptions{
		TestName:          t.Name(),
		Client:            client,
		NetworkID:         network,
		BlockDatabaseFile: interchaintest.DefaultBlockDatabaseFilepath(),
		SkipPathCreation:  true, // Skip path creation, so we can have granular control over the process
	}))

	t.Cleanup(func() {
		_ = ic.Close()
	})

	// Create and Fund User Wallets
	fundAmount := int64(10_000_000_000)
	users := interchaintest.GetAndFundTestUsers(t, ctx, "default", int64(fundAmount), quicksilverd, juno)
	quicksilverd1User := users[0]
	juno1User := users[1]

	err = testutil.WaitForBlocks(ctx, 10, quicksilverd, juno)
	require.NoError(t, err)

	quicksilverd1UserBalInitial, err := quicksilverd.GetBalance(ctx, quicksilverd1User.FormattedAddress(), quicksilverd.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, fundAmount, quicksilverd1UserBalInitial)

	juno1UserBalInitial, err := juno.GetBalance(ctx, juno1User.FormattedAddress(), juno.Config().Denom)
	require.NoError(t, err)
	require.Equal(t, fundAmount, juno1UserBalInitial)

	// Generate a new IBC path
	err = r.GeneratePath(ctx, eRep, quicksilverd.Config().ChainID, juno.Config().ChainID, pathQuicksilverJuno)
	require.NoError(t, err)

	// Create a new connection
	err = r.CreateConnections(ctx, eRep, pathQuicksilverJuno)
	require.NoError(t, err)

	// Create a new channel
	err = r.CreateChannel(ctx, eRep, pathQuicksilverJuno, ibc.DefaultChannelOpts())
	require.NoError(t, err)

	// Query for the newly created channel
	quickChannels, err := r.GetChannels(ctx, eRep, quicksilverd.Config().ChainID)
	require.NoError(t, err)
	fmt.Println("Channel", quickChannels)

	// Start the relayer and set the cleanup function.
	require.NoError(t, r.StartRelayer(ctx, eRep, pathQuicksilverJuno))
	t.Cleanup(
		func() {
			err := r.StopRelayer(ctx, eRep)
			if err != nil {
				panic(fmt.Errorf("an error occurred while stopping the relayer: %s", err))
			}
		},
	)

	proposal := cosmos.TxProposalv1{
		Metadata: "none",
		Deposit:  "500000000" + quicksilverd.Config().Denom, // greater than min deposit
		Title:    "title",
		Summary:  "suma",
	}

	content := istypes.RegisterZoneProposal{
		Title:            "register lstest-1 zone",
		Description:      "register lstest-1 zone with multisend and lsm enabled",
		ConnectionId:     "connection-0",
		BaseDenom:        "uatom",
		LocalDenom:       "uqatom",
		AccountPrefix:    "cosmos",
		DepositsEnabled:  true,
		UnbondingEnabled: true,
		LiquidityModule:  false,
		ReturnToSender:   true,
		Decimals:         6,
	}

	check, err := cdctypes.NewAnyWithValue(&content)

	message := govv1.MsgExecLegacyContent{
		Content:   check,
		Authority: "quick10d07y265gmmuvt4z0w9aw880jnsr700j3xrh0p",
	}
	msg, err := quicksilverd.Config().EncodingConfig.Codec.MarshalInterfaceJSON(&message)
	fmt.Println("Msg: ", string(msg))
	require.NoError(t, err)
	proposal.Messages = append(proposal.Messages, msg)

	// Submit Proposal
	proposalTx, err := quicksilverd.SubmitProposal(ctx, quicksilverd1User.KeyName(), proposal)
	require.NoError(t, err, "error submitting proposal tx")

	height, err := quicksilverd.Height(ctx)
	require.NoError(t, err, "error fetching height before submit upgrade proposal")

	err = quicksilverd.VoteOnProposalAllValidators(ctx, proposalTx.ProposalID, cosmos.ProposalVoteYes)
	require.NoError(t, err, "failed to submit votes")

	_, err = cosmos.PollForProposalStatus(ctx, quicksilverd, height, height+heightDelta, proposalTx.ProposalID, cosmos.ProposalStatusPassed)
	require.NoError(t, err, "proposal status did not change to passed in expected number of blocks")

	err = testutil.WaitForBlocks(ctx, 2, quicksilverd)
	require.NoError(t, err)

	stdout, _, err := quicksilverd.Validators[0].ExecQuery(ctx, "interchainstaking", "zones", "--output=json")

	t.Logf("zones: %s", stdout)
	// TODO: Need address not nil
	require.NoError(t, err)
}

func modifyGenesisShortProposals(votingPeriod, maxDepositPeriod string) func(ibc.ChainConfig, []byte) ([]byte, error) {
	return func(chainConfig ibc.ChainConfig, genbz []byte) ([]byte, error) {
		g := make(map[string]interface{})
		if err := json.Unmarshal(genbz, &g); err != nil {
			return nil, fmt.Errorf("failed to unmarshal genesis file: %w", err)
		}
		if err := dyno.Set(g, votingPeriod, "app_state", "gov", "params", "voting_period"); err != nil {
			return nil, fmt.Errorf("failed to set voting period in genesis json: %w", err)
		}
		if err := dyno.Set(g, maxDepositPeriod, "app_state", "gov", "params", "max_deposit_period"); err != nil {
			return nil, fmt.Errorf("failed to set voting period in genesis json: %w", err)
		}
		if err := dyno.Set(g, chainConfig.Denom, "app_state", "gov", "params", "min_deposit", 0, "denom"); err != nil {
			return nil, fmt.Errorf("failed to set voting period in genesis json: %w", err)
		}
		out, err := json.Marshal(g)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal genesis bytes to json: %w", err)
		}
		return out, nil
	}
}
