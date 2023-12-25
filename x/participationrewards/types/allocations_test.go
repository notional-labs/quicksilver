package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"
	sdkmath "cosmossdk.io/math"

	"github.com/quicksilver-zone/quicksilver/x/participationrewards/types"
)

func TestGetRewardsAllocations(t *testing.T) {
	type args struct {
		moduleBalance math.Int
		proportions   types.DistributionProportions
	}
	tests := []struct {
		name    string
		args    args
		want    *types.RewardsAllocation
		wantErr string
	}{
		{
			"empty_params",
			args{},
			nil,
			"balance is zero, nothing to allocate",
		},
		{
			"invalid_no_balance",
			args{
				sdkmath.NewInt(0),
				types.DistributionProportions{
					ValidatorSelectionAllocation: sdkmath.LegacyMustNewDecFromStr("0.34"),
					HoldingsAllocation:           sdkmath.LegacyMustNewDecFromStr("0.33"),
					LockupAllocation:             sdkmath.LegacyMustNewDecFromStr("0.33"),
				},
			},
			nil,
			"balance is zero, nothing to allocate",
		},
		{
			"invalid_proportions_gt",
			args{
				sdkmath.NewInt(1000000000),
				types.DistributionProportions{
					ValidatorSelectionAllocation: sdkmath.LegacyMustNewDecFromStr("0.5"),
					HoldingsAllocation:           sdkmath.LegacyMustNewDecFromStr("0.5"),
					LockupAllocation:             sdkmath.LegacyMustNewDecFromStr("0.5"),
				},
			},
			nil,
			"total distribution proportions must be 1.0: got 1.50",
		},
		{
			"invalid_proportions_lt",
			args{
				sdkmath.NewInt(1000000000),
				types.DistributionProportions{
					ValidatorSelectionAllocation: sdkmath.LegacyMustNewDecFromStr("0.3"),
					HoldingsAllocation:           sdkmath.LegacyMustNewDecFromStr("0.3"),
					LockupAllocation:             sdkmath.LegacyMustNewDecFromStr("0.3"),
				},
			},
			nil,
			"total distribution proportions must be 1.0: got 0.90",
		},
		{
			"valid",
			args{
				sdkmath.NewInt(1000000000),
				types.DistributionProportions{
					ValidatorSelectionAllocation: sdkmath.LegacyMustNewDecFromStr("0.34"),
					HoldingsAllocation:           sdkmath.LegacyMustNewDecFromStr("0.33"),
					LockupAllocation:             sdkmath.LegacyMustNewDecFromStr("0.33"),
				},
			},
			&types.RewardsAllocation{
				ValidatorSelection: sdkmath.NewInt(340000000),
				Holdings:           sdkmath.NewInt(330000000),
				Lockup:             sdkmath.NewInt(330000000),
			},
			"",
		},
		{
			"valid",
			args{
				sdkmath.NewInt(1000000000),
				types.DistributionProportions{
					ValidatorSelectionAllocation: sdkmath.LegacyMustNewDecFromStr("0.5"),
					HoldingsAllocation:           sdkmath.LegacyMustNewDecFromStr("0.25"),
					LockupAllocation:             sdkmath.LegacyMustNewDecFromStr("0.25"),
				},
			},
			&types.RewardsAllocation{
				ValidatorSelection: sdkmath.NewInt(500000000),
				Holdings:           sdkmath.NewInt(250000000),
				Lockup:             sdkmath.NewInt(250000000),
			},
			"",
		},
		{
			"valid",
			args{
				sdkmath.NewInt(1000000000),
				types.DistributionProportions{
					ValidatorSelectionAllocation: sdkmath.LegacyMustNewDecFromStr("0.6"),
					HoldingsAllocation:           sdkmath.LegacyMustNewDecFromStr("0.4"),
					LockupAllocation:             sdkmath.LegacyMustNewDecFromStr("0"),
				},
			},
			&types.RewardsAllocation{
				ValidatorSelection: sdkmath.NewInt(600000000),
				Holdings:           sdkmath.NewInt(400000000),
				Lockup:             sdkmath.NewInt(0),
			},
			"",
		},
		{
			"valid",
			args{
				sdkmath.NewInt(164133471813),
				types.DistributionProportions{
					ValidatorSelectionAllocation: sdkmath.LegacyMustNewDecFromStr("0.34"),
					HoldingsAllocation:           sdkmath.LegacyMustNewDecFromStr("0.33"),
					LockupAllocation:             sdkmath.LegacyMustNewDecFromStr("0.33"),
				},
			},
			&types.RewardsAllocation{
				ValidatorSelection: sdkmath.NewInt(55805380417),
				Holdings:           sdkmath.NewInt(54164045698),
				Lockup:             sdkmath.NewInt(54164045698),
			},
			"",
		},
		{
			"valid",
			args{
				sdkmath.NewInt(164133471813),
				types.DistributionProportions{
					ValidatorSelectionAllocation: sdkmath.LegacyMustNewDecFromStr("0.5"),
					HoldingsAllocation:           sdkmath.LegacyMustNewDecFromStr("0.25"),
					LockupAllocation:             sdkmath.LegacyMustNewDecFromStr("0.25"),
				},
			},
			&types.RewardsAllocation{
				ValidatorSelection: sdkmath.NewInt(82066735907),
				Holdings:           sdkmath.NewInt(41033367953),
				Lockup:             sdkmath.NewInt(41033367953),
			},
			"",
		},
		{
			"valid",
			args{
				sdkmath.NewInt(164133471813),
				types.DistributionProportions{
					ValidatorSelectionAllocation: sdkmath.LegacyMustNewDecFromStr("0.6"),
					HoldingsAllocation:           sdkmath.LegacyMustNewDecFromStr("0.4"),
					LockupAllocation:             sdkmath.LegacyMustNewDecFromStr("0"),
				},
			},
			&types.RewardsAllocation{
				ValidatorSelection: sdkmath.NewInt(98480083088),
				Holdings:           sdkmath.NewInt(65653388725),
				Lockup:             sdkmath.NewInt(0),
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := types.GetRewardsAllocations(tt.args.moduleBalance, tt.args.proportions)
			if tt.wantErr != "" {
				require.Error(t, err)
				require.Nil(t, got)
				require.Contains(t, err.Error(), tt.wantErr)
				t.Logf("Error: %v", err)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, got)
			require.Equal(t, tt.want, got)
		})
	}
}
