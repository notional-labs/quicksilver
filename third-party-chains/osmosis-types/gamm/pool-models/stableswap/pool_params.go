package stableswap

import (
	sdkmath "cosmossdk.io/math"

	"github.com/quicksilver-zone/quicksilver/third-party-chains/osmosis-types/gamm"
)

func (params PoolParams) Validate() error {
	if params.ExitFee.IsNegative() {
		return gamm.ErrNegativeExitFee
	}

	if params.ExitFee.GTE(sdkmath.LegacyOneDec()) {
		return gamm.ErrTooMuchExitFee
	}

	if params.SwapFee.IsNegative() {
		return gamm.ErrNegativeSwapFee
	}

	if params.SwapFee.GTE(sdkmath.LegacyOneDec()) {
		return gamm.ErrTooMuchSwapFee
	}
	return nil
}
