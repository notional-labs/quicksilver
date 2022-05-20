package keeper

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	tmclienttypes "github.com/cosmos/ibc-go/v3/modules/light-clients/07-tendermint/types"

	commitmenttypes "github.com/cosmos/ibc-go/v3/modules/core/23-commitment/types"
	"github.com/ingenuity-build/quicksilver/x/interchainquery/types"
)

type msgServer struct {
	*Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: &keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) SubmitQueryResponse(goCtx context.Context, msg *types.MsgSubmitQueryResponse) (*types.MsgSubmitQueryResponseResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	q, found := k.GetQuery(ctx, msg.QueryId)
	if found {
		pathParts := strings.Split(q.QueryType, "/")
		if pathParts[len(pathParts)-1] == "key" {
			if msg.ProofOps == nil {
				k.Logger(ctx).Error("KV lookup requires a proof")
				return nil, fmt.Errorf("unable to validate proof. No proof submitted")
			}
			connection, _ := k.IBCKeeper.ConnectionKeeper.GetConnection(ctx, q.ConnectionId)

			height := clienttypes.NewHeight(0, uint64(msg.Height))
			consensusState, _ := k.IBCKeeper.ClientKeeper.GetClientConsensusState(ctx, connection.ClientId, height)
			clientState, _ := k.IBCKeeper.ClientKeeper.GetClientState(ctx, connection.ClientId)
			path := commitmenttypes.NewMerklePath("/" + q.QueryType)
			merkleProof, _ := commitmenttypes.ConvertProofs(msg.ProofOps)

			tmclientstate, _ := clientState.(*tmclienttypes.ClientState)
			if err := merkleProof.VerifyMembership(tmclientstate.ProofSpecs, consensusState.GetRoot(), path, msg.Result); err != nil {
				k.Logger(ctx).Error("Unable to verify proof")
				return nil, fmt.Errorf("unable to verify proof: %s", err)
			}
		}
		for _, module := range k.callbacks {
			if module.Has(msg.QueryId) {
				err := module.Call(ctx, msg.QueryId, msg.Result, q)
				if err != nil {
					k.Logger(ctx).Error("Error in callback", "error", err, "msg", msg.QueryId, "result", msg.Result, "type", q.QueryType, "params", q.Request)
					return nil, err
				}
			}
		}
		//q.LastHeight = sdk.NewInt(ctx.BlockHeight())

		if err := k.SetDatapointForId(ctx, msg.QueryId, msg.Result, sdk.NewInt(msg.Height)); err != nil {
			return nil, err
		}

		if q.Period.IsNegative() {
			k.DeleteQuery(ctx, msg.QueryId)
		} else {
			k.SetQuery(ctx, q)
		}

	} else {
		return &types.MsgSubmitQueryResponseResponse{}, nil // technically this is an error, but will cause the entire tx to fail if we have one 'bad' message, so we can just no-op here.
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})

	return &types.MsgSubmitQueryResponseResponse{}, nil
}
