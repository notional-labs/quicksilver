package types

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	"gopkg.in/yaml.v2"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyDistributionProportions = []byte("DistributionProportions")
	KeyClaimsEnabled           = []byte("ClaimsEnabled")

	DefaultValidatorSelectionAllocation = sdkmath.LegacyNewDecWithPrec(34, 2)
	DefaultHoldingsAllocation           = sdkmath.LegacyNewDecWithPrec(33, 2)
	DefaultLockupAllocation             = sdkmath.LegacyNewDecWithPrec(33, 2)
	DefaultClaimsEnabled                = false
)

// ParamKeyTable for participationrewards module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new ics Params instance.
func NewParams(
	validatorSelectionAllocation sdkmath.LegacyDec,
	holdingsAllocation sdkmath.LegacyDec,
	lockupAllocation sdkmath.LegacyDec,
	claimsEnabled bool,
) Params {
	return Params{
		DistributionProportions: DistributionProportions{
			ValidatorSelectionAllocation: validatorSelectionAllocation,
			HoldingsAllocation:           holdingsAllocation,
			LockupAllocation:             lockupAllocation,
		},
		ClaimsEnabled: claimsEnabled,
	}
}

// DefaultParams default ics params.
func DefaultParams() Params {
	return NewParams(
		DefaultValidatorSelectionAllocation,
		DefaultHoldingsAllocation,
		DefaultLockupAllocation,
		DefaultClaimsEnabled,
	)
}

// ParamSetPairs implements params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDistributionProportions, &p.DistributionProportions, validateDistributionProportions),
		paramtypes.NewParamSetPair(KeyClaimsEnabled, &p.ClaimsEnabled, validateBoolean),
	}
}

func validateDistributionProportions(i interface{}) error {
	dp, ok := i.(DistributionProportions)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return dp.ValidateBasic()
}

func validateBoolean(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

// Validate performs stateless validity checks on params.
func (p *Params) Validate() error {
	return validateDistributionProportions(p.DistributionProportions)
}

// String implements the Stringer interface.
func (p *Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}
