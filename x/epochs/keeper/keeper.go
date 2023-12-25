package keeper

import (
	"fmt"

	"cosmossdk.io/log"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/quicksilver-zone/quicksilver/x/epochs/types"
)

// Keeper of this module maintains collections of epochs and hooks.
type Keeper struct {
	cdc      codec.Codec
	storeKey storetypes.StoreKey
	hooks    types.EpochHooks
}

// NewKeeper returns a new instance of epochs Keeper.
func NewKeeper(cdc codec.Codec, storeKey storetypes.StoreKey) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

// SetHooks set the epoch hooks.
func (k *Keeper) SetHooks(eh types.EpochHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set epochs hooks twice")
	}

	k.hooks = eh

	return k
}

// Logger returns a module-specific logger.
func (*Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
