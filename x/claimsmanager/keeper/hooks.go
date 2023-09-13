package keeper

import (
	epochstypes "github.com/quicksilver-zone/quicksilver/x/epochs/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) error {
	if epochIdentifier == epochstypes.EpochIdentifierEpoch && epochNumber > 1 {
		if err := k.StoreSelfConsensusState(ctx, "epoch"); err != nil {
			k.Logger(ctx).Error("unable to store consensus state", "error", err)
			return err
		}
	}
	return nil
}

func (k Keeper) AfterEpochEnd(_ sdk.Context, _ string, _ int64) error {
	return nil
}

// ___________________________________________________________________________________________________

// Hooks wrapper struct for incentives keeper.
type Hooks struct {
	k Keeper
}

var _ epochstypes.EpochHooks = Hooks{}

func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// epochs hooks.
func (h Hooks) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) error {
	return h.k.BeforeEpochStart(ctx, epochIdentifier, epochNumber)
}

func (h Hooks) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) error {
	return h.k.AfterEpochEnd(ctx, epochIdentifier, epochNumber)
}
