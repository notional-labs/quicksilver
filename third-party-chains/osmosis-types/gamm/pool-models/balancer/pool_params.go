package balancer

import (
	"errors"

	sdkmath "cosmossdk.io/math"

	"github.com/quicksilver-zone/quicksilver/third-party-chains/osmosis-types/gamm"
)

func NewPoolParams(swapFee, exitFee sdkmath.LegacyDec, params *SmoothWeightChangeParams) PoolParams {
	return PoolParams{
		SwapFee:                  swapFee,
		ExitFee:                  exitFee,
		SmoothWeightChangeParams: params,
	}
}

func (params PoolParams) Validate(poolWeights []PoolAsset) error {
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

	if params.SmoothWeightChangeParams != nil {
		targetWeights := params.SmoothWeightChangeParams.TargetPoolWeights
		// Ensure it has the right number of weights
		if len(targetWeights) != len(poolWeights) {
			return gamm.ErrPoolParamsInvalidNumDenoms
		}
		// Validate all user specified weights
		for _, v := range targetWeights {
			err := ValidateUserSpecifiedWeight(v.Weight)
			if err != nil {
				return err
			}
		}
		// Ensure that all the target weight denoms are same as pool asset weights
		sortedTargetPoolWeights := sortPoolAssetsOutOfPlaceByDenom(targetWeights)
		sortedPoolWeights := sortPoolAssetsOutOfPlaceByDenom(poolWeights)
		for i, v := range sortedPoolWeights {
			if sortedTargetPoolWeights[i].Token.Denom != v.Token.Denom {
				return gamm.ErrPoolParamsInvalidDenom
			}
		}

		// No start time validation needed

		// We do not need to validate InitialPoolWeights, as we set that ourselves
		// in setInitialPoolParams

		// TODO: Is there anything else we can validate for duration?
		if params.SmoothWeightChangeParams.Duration <= 0 {
			return errors.New("params.SmoothWeightChangeParams must have a positive duration")
		}
	}

	return nil
}

func (params PoolParams) GetPoolSwapFee() sdkmath.LegacyDec {
	return params.SwapFee
}

func (params PoolParams) GetPoolExitFee() sdkmath.LegacyDec {
	return params.ExitFee
}
