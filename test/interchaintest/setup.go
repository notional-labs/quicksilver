package interchaintest

import (
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	istypes "github.com/ingenuity-build/quicksilver/x/interchainstaking/types"
	"github.com/strangelove-ventures/interchaintest/v7/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v7/ibc"
)

var (
	QuickSilverE2ERepo  = "quicksilverzone/quicksilver-e2e"
	QuicksilverMainRepo = "quicksilverzone/quicksilver"

	repo, version = GetDockerImageInfo()

	QuicksilverImage = ibc.DockerImage{
		Repository: repo,
		Version:    version,
		UidGid:     "1025:1025",
	}

	IcaImage = ibc.DockerImage{
		Repository: "quicksilverzone/testzone",
		Version:    "latest",
		UidGid:     "1025:1025",
	}

	XccLookupImage = ibc.DockerImage{
		Repository: "quicksilverzone/xcclookup",
		Version:    "v0.4.3",
		UidGid:     "1026:1026",
	}

	ICQImage = ibc.DockerImage{
		Repository: "quicksilverzone/interchain-queries",
		Version:    "e2e",
		UidGid:     "1027:1027",
	}

	pathQuicksilverJuno = "quicksilver-juno"
	genesisWalletAmount = int64(10_000_000)
)

func createConfig() (ibc.ChainConfig, error) {
	return ibc.ChainConfig{
			Type:                "cosmos",
			Name:                "quicksilver",
			ChainID:             "quicksilver-2",
			Images:              []ibc.DockerImage{QuicksilverImage},
			Bin:                 "quicksilverd",
			Bech32Prefix:        "quick",
			Denom:               "uqck",
			GasPrices:           "0.0uqck",
			GasAdjustment:       1.1,
			TrustingPeriod:      "112h",
			NoHostMount:         false,
			ModifyGenesis:       nil,
			ConfigFileOverrides: nil,
			EncodingConfig:      nil,
			// TODO: Need add SideCar to interchain-test v7

			// SidecarConfigs: []ibc.SidecarConfig{
			// 	{
			// 		ProcessName:      "icq",
			// 		Image:            ICQImage,
			// 		Ports:            []string{"2112"},
			// 		StartCmd:         []string{"interchain-queries", "run", "--home", "/var/sidecar-processes/icq"},
			// 		PreStart:         true,
			// 		ValidatorProcess: false,
			// 	},
			// 	{
			// 		ProcessName:      "xcc",
			// 		Image:            XccLookupImage,
			// 		Ports:            []string{"3033"},
			// 		StartCmd:         []string{"/xcc", "-a", "serve", "-f", "/var/sidecar/processes/xcc/config.yaml"},
			// 		PreStart:         true,
			// 		ValidatorProcess: false,
			// 	},
			// },
		},
		nil
}

// quicksilverEncoding registers the Quicksilver specific module codecs so that the associated types and msgs
// will be supported when writing to the blocksdb sqlite database.
func quicksilverEncoding() *testutil.TestEncodingConfig {
	cfg := cosmos.DefaultEncoding()

	// register custom types
	istypes.RegisterInterfaces(cfg.InterfaceRegistry)
	govv1.RegisterInterfaces(cfg.InterfaceRegistry)
	return &cfg
}
