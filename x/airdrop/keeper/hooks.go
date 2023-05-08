package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	epochstypes "github.com/ingenuity-build/quicksilver/x/epochs/types"
)

<<<<<<< HEAD
func (k Keeper) BeforeEpochStart(_ sdk.Context, _ string, _ int64) {
}

func (k Keeper) AfterEpochEnd(_ sdk.Context, _ string, _ int64) {
=======
func (k *Keeper) BeforeEpochStart(_ sdk.Context, _ string, _ int64) error {
	return nil
}

func (k *Keeper) AfterEpochEnd(_ sdk.Context, _ string, _ int64) error {
	return nil
>>>>>>> origin/develop
}

// ___________________________________________________________________________________________________

// Hooks wrapper struct for airdrop keeper.
type Hooks struct {
	k *Keeper
}

var _ epochstypes.EpochHooks = Hooks{}

func (k *Keeper) Hooks() Hooks {
	return Hooks{k}
}

// epochs hooks.

func (h Hooks) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) error {
	return h.k.BeforeEpochStart(ctx, epochIdentifier, epochNumber)
}

func (h Hooks) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) error {
	return h.k.AfterEpochEnd(ctx, epochIdentifier, epochNumber)
}
