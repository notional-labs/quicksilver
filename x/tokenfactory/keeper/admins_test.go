package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/ingenuity-build/quicksilver/x/tokenfactory/types"
)

func (s *KeeperTestSuite) TestAdminMsgs() {
	addr0bal := int64(0)
	addr1bal := int64(0)

	bankKeeper := s.App.BankKeeper

	s.CreateDefaultDenom()
	// Make sure that the admin is set correctly
	queryRes, err := s.queryClient.DenomAuthorityMetadata(s.Ctx.Context(), &types.QueryDenomAuthorityMetadataRequest{
		Denom: s.defaultDenom,
	})
	s.Require().NoError(err)
	s.Require().Equal(s.TestAccs[0].String(), queryRes.AuthorityMetadata.Admin)

	// Test minting to admins own account
	_, err = s.msgServer.Mint(sdk.WrapSDKContext(s.Ctx), types.NewMsgMint(s.TestAccs[0].String(), sdk.NewInt64Coin(s.defaultDenom, 10)))
	addr0bal += 10
	s.Require().NoError(err)
	s.Require().True(bankKeeper.GetBalance(s.Ctx, s.TestAccs[0], s.defaultDenom).Amount.Int64() == addr0bal, bankKeeper.GetBalance(s.Ctx, s.TestAccs[0], s.defaultDenom))

	// // Test force transferring
	// _, err = suite.msgServer.ForceTransfer(sdk.WrapSDKContext(suite.Ctx), types.NewMsgForceTransfer(suite.TestAccs[0].String(), sdk.NewInt64Coin(denom, 5), suite.TestAccs[1].String(), suite.TestAccs[0].String()))
	// suite.Require().NoError(err)
	// suite.Require().True(bankKeeper.GetBalance(suite.Ctx, suite.TestAccs[0], denom).IsEqual(sdk.NewInt64Coin(denom, 15)))
	// suite.Require().True(bankKeeper.GetBalance(suite.Ctx, suite.TestAccs[1], denom).IsEqual(sdk.NewInt64Coin(denom, 5)))

	// Test burning from own account
	_, err = s.msgServer.Burn(sdk.WrapSDKContext(s.Ctx), types.NewMsgBurn(s.TestAccs[0].String(), sdk.NewInt64Coin(s.defaultDenom, 5)))
	s.Require().NoError(err)
	s.Require().True(bankKeeper.GetBalance(s.Ctx, s.TestAccs[1], s.defaultDenom).Amount.Int64() == addr1bal)

	// Test Change Admin
	_, err = s.msgServer.ChangeAdmin(sdk.WrapSDKContext(s.Ctx), types.NewMsgChangeAdmin(s.TestAccs[0].String(), s.defaultDenom, s.TestAccs[1].String()))
	s.Require().NoError(err)
	queryRes, err = s.queryClient.DenomAuthorityMetadata(s.Ctx.Context(), &types.QueryDenomAuthorityMetadataRequest{
		Denom: s.defaultDenom,
	})
	s.Require().NoError(err)
	s.Require().Equal(s.TestAccs[1].String(), queryRes.AuthorityMetadata.Admin)

	// Make sure old admin can no longer do actions
	_, err = s.msgServer.Burn(sdk.WrapSDKContext(s.Ctx), types.NewMsgBurn(s.TestAccs[0].String(), sdk.NewInt64Coin(s.defaultDenom, 5)))
	s.Require().Error(err)

	// Make sure the new admin works
	_, err = s.msgServer.Mint(sdk.WrapSDKContext(s.Ctx), types.NewMsgMint(s.TestAccs[1].String(), sdk.NewInt64Coin(s.defaultDenom, 5)))
	addr1bal += 5
	s.Require().NoError(err)
	s.Require().True(bankKeeper.GetBalance(s.Ctx, s.TestAccs[1], s.defaultDenom).Amount.Int64() == addr1bal)

	// Try setting admin to empty
	_, err = s.msgServer.ChangeAdmin(sdk.WrapSDKContext(s.Ctx), types.NewMsgChangeAdmin(s.TestAccs[1].String(), s.defaultDenom, ""))
	s.Require().NoError(err)
	queryRes, err = s.queryClient.DenomAuthorityMetadata(s.Ctx.Context(), &types.QueryDenomAuthorityMetadataRequest{
		Denom: s.defaultDenom,
	})
	s.Require().NoError(err)
	s.Require().Equal("", queryRes.AuthorityMetadata.Admin)
}

// TestMintDenom ensures the following properties of the MintMessage:
// * Noone can mint tokens for a denom that doesn't exist.
// * Only the admin of a denom can mint tokens for it.
// * The admin of a denom can mint tokens for it.
func (s *KeeperTestSuite) TestMintDenom() {
	var addr0bal int64

	// Create a denom
	s.CreateDefaultDenom()

	for _, tc := range []struct {
		desc      string
		amount    int64
		mintDenom string
		admin     string
		valid     bool
	}{
		{
			desc:      "denom does not exist",
			amount:    10,
			mintDenom: "factory/osmo1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44/evmos",
			admin:     s.TestAccs[0].String(),
			valid:     false,
		},
		{
			desc:      "mint is not by the admin",
			amount:    10,
			mintDenom: s.defaultDenom,
			admin:     s.TestAccs[1].String(),
			valid:     false,
		},
		{
			desc:      "success case",
			amount:    10,
			mintDenom: s.defaultDenom,
			admin:     s.TestAccs[0].String(),
			valid:     true,
		},
	} {
		s.Run(fmt.Sprintf("Case %s", tc.desc), func() {
			// Test minting to admins own account
			bankKeeper := s.App.BankKeeper
			_, err := s.msgServer.Mint(sdk.WrapSDKContext(s.Ctx), types.NewMsgMint(tc.admin, sdk.NewInt64Coin(tc.mintDenom, 10)))
			if tc.valid {
				addr0bal += 10
				s.Require().NoError(err)
				s.Require().Equal(bankKeeper.GetBalance(s.Ctx, s.TestAccs[0], s.defaultDenom).Amount.Int64(), addr0bal, bankKeeper.GetBalance(s.Ctx, s.TestAccs[0], s.defaultDenom))
			} else {
				s.Require().Error(err)
			}
		})
	}
}

func (s *KeeperTestSuite) TestBurnDenom() {
	var addr0bal int64

	// Create a denom.
	s.CreateDefaultDenom()

	// mint 10 default token for testAcc[0]
	_, err := s.msgServer.Mint(sdk.WrapSDKContext(s.Ctx), types.NewMsgMint(s.TestAccs[0].String(), sdk.NewInt64Coin(s.defaultDenom, 10)))
	s.Require().NoError(err)

	addr0bal += 10

	for _, tc := range []struct {
		desc      string
		amount    int64
		burnDenom string
		admin     string
		valid     bool
	}{
		{
			desc:      "denom does not exist",
			amount:    10,
			burnDenom: "factory/osmo1t7egva48prqmzl59x5ngv4zx0dtrwewc9m7z44/evmos",
			admin:     s.TestAccs[0].String(),
			valid:     false,
		},
		{
			desc:      "burn is not by the admin",
			amount:    10,
			burnDenom: s.defaultDenom,
			admin:     s.TestAccs[1].String(),
			valid:     false,
		},
		{
			desc:      "burn amount is bigger than minted amount",
			amount:    1000,
			burnDenom: s.defaultDenom,
			admin:     s.TestAccs[1].String(),
			valid:     false,
		},
		{
			desc:      "success case",
			amount:    10,
			burnDenom: s.defaultDenom,
			admin:     s.TestAccs[0].String(),
			valid:     true,
		},
	} {
		s.Run(fmt.Sprintf("Case %s", tc.desc), func() {
			// Test minting to admins own account
			bankKeeper := s.App.BankKeeper
			_, err := s.msgServer.Burn(sdk.WrapSDKContext(s.Ctx), types.NewMsgBurn(tc.admin, sdk.NewInt64Coin(tc.burnDenom, 10)))
			if tc.valid {
				addr0bal -= 10
				s.Require().NoError(err)
				s.Require().True(bankKeeper.GetBalance(s.Ctx, s.TestAccs[0], s.defaultDenom).Amount.Int64() == addr0bal, bankKeeper.GetBalance(s.Ctx, s.TestAccs[0], s.defaultDenom))
			} else {
				s.Require().Error(err)
				s.Require().True(bankKeeper.GetBalance(s.Ctx, s.TestAccs[0], s.defaultDenom).Amount.Int64() == addr0bal, bankKeeper.GetBalance(s.Ctx, s.TestAccs[0], s.defaultDenom))
			}
		})
	}
}

func (s *KeeperTestSuite) TestChangeAdminDenom() {
	for _, tc := range []struct {
		desc                    string
		msgChangeAdmin          func(denom string) *types.MsgChangeAdmin
		expectedChangeAdminPass bool
		expectedAdminIndex      int
		msgMint                 func(denom string) *types.MsgMint
		expectedMintPass        bool
	}{
		{
			desc: "creator admin can't mint after setting to '' ",
			msgChangeAdmin: func(denom string) *types.MsgChangeAdmin {
				return types.NewMsgChangeAdmin(s.TestAccs[0].String(), denom, "")
			},
			expectedChangeAdminPass: true,
			expectedAdminIndex:      -1,
			msgMint: func(denom string) *types.MsgMint {
				return types.NewMsgMint(s.TestAccs[0].String(), sdk.NewInt64Coin(denom, 5))
			},
			expectedMintPass: false,
		},
		{
			desc: "non-admins can't change the existing admin",
			msgChangeAdmin: func(denom string) *types.MsgChangeAdmin {
				return types.NewMsgChangeAdmin(s.TestAccs[1].String(), denom, s.TestAccs[2].String())
			},
			expectedChangeAdminPass: false,
			expectedAdminIndex:      0,
		},
		{
			desc: "success change admin",
			msgChangeAdmin: func(denom string) *types.MsgChangeAdmin {
				return types.NewMsgChangeAdmin(s.TestAccs[0].String(), denom, s.TestAccs[1].String())
			},
			expectedAdminIndex:      1,
			expectedChangeAdminPass: true,
			msgMint: func(denom string) *types.MsgMint {
				return types.NewMsgMint(s.TestAccs[1].String(), sdk.NewInt64Coin(denom, 5))
			},
			expectedMintPass: true,
		},
	} {
		s.Run(fmt.Sprintf("Case %s", tc.desc), func() {
			// setup test
			s.SetupTest()

			// Create a denom and mint
			res, err := s.msgServer.CreateDenom(sdk.WrapSDKContext(s.Ctx), types.NewMsgCreateDenom(s.TestAccs[0].String(), "bitcoin"))
			s.Require().NoError(err)

			testDenom := res.GetNewTokenDenom()

			_, err = s.msgServer.Mint(sdk.WrapSDKContext(s.Ctx), types.NewMsgMint(s.TestAccs[0].String(), sdk.NewInt64Coin(testDenom, 10)))
			s.Require().NoError(err)

			_, err = s.msgServer.ChangeAdmin(sdk.WrapSDKContext(s.Ctx), tc.msgChangeAdmin(testDenom))
			if tc.expectedChangeAdminPass {
				s.Require().NoError(err)
			} else {
				s.Require().Error(err)
			}

			queryRes, err := s.queryClient.DenomAuthorityMetadata(s.Ctx.Context(), &types.QueryDenomAuthorityMetadataRequest{
				Denom: testDenom,
			})
			s.Require().NoError(err)

			// expectedAdminIndex with negative value is assumed as admin with value of ""
			const emptyStringAdminIndexFlag = -1
			if tc.expectedAdminIndex == emptyStringAdminIndexFlag {
				s.Require().Equal("", queryRes.AuthorityMetadata.Admin)
			} else {
				s.Require().Equal(s.TestAccs[tc.expectedAdminIndex].String(), queryRes.AuthorityMetadata.Admin)
			}

			// we test mint to test if admin authority is performed properly after admin change.
			if tc.msgMint != nil {
				_, err := s.msgServer.Mint(sdk.WrapSDKContext(s.Ctx), tc.msgMint(testDenom))
				if tc.expectedMintPass {
					s.Require().NoError(err)
				} else {
					s.Require().Error(err)
				}
			}
		})
	}
}

func (s *KeeperTestSuite) TestSetDenomMetaData() {
	// setup test
	s.SetupTest()
	s.CreateDefaultDenom()

	for _, tc := range []struct {
		desc                string
		msgSetDenomMetadata types.MsgSetDenomMetadata
		expectedPass        bool
	}{
		{
			desc: "successful set denom metadata",
			msgSetDenomMetadata: *types.NewMsgSetDenomMetadata(s.TestAccs[0].String(), banktypes.Metadata{
				Description: "yeehaw",
				DenomUnits: []*banktypes.DenomUnit{
					{
						Denom:    s.defaultDenom,
						Exponent: 0,
					},
					{
						Denom:    "uqck",
						Exponent: 6,
					},
				},
				Base:    s.defaultDenom,
				Display: "uqck",
				Name:    "OSMO",
				Symbol:  "OSMO",
			}),
			expectedPass: true,
		},
		{
			desc: "non existent factory denom name",
			msgSetDenomMetadata: *types.NewMsgSetDenomMetadata(s.TestAccs[0].String(), banktypes.Metadata{
				Description: "yeehaw",
				DenomUnits: []*banktypes.DenomUnit{
					{
						Denom:    fmt.Sprintf("factory/%s/litecoin", s.TestAccs[0].String()),
						Exponent: 0,
					},
					{
						Denom:    "uqck",
						Exponent: 6,
					},
				},
				Base:    fmt.Sprintf("factory/%s/litecoin", s.TestAccs[0].String()),
				Display: "uqck",
				Name:    "OSMO",
				Symbol:  "OSMO",
			}),
			expectedPass: false,
		},
		{
			desc: "non-factory denom",
			msgSetDenomMetadata: *types.NewMsgSetDenomMetadata(s.TestAccs[0].String(), banktypes.Metadata{
				Description: "yeehaw",
				DenomUnits: []*banktypes.DenomUnit{
					{
						Denom:    "uqck",
						Exponent: 0,
					},
					{
						Denom:    "uqcko",
						Exponent: 6,
					},
				},
				Base:    "uqck",
				Display: "uqcko",
				Name:    "OSMO",
				Symbol:  "OSMO",
			}),
			expectedPass: false,
		},
		{
			desc: "wrong admin",
			msgSetDenomMetadata: *types.NewMsgSetDenomMetadata(s.TestAccs[1].String(), banktypes.Metadata{
				Description: "yeehaw",
				DenomUnits: []*banktypes.DenomUnit{
					{
						Denom:    s.defaultDenom,
						Exponent: 0,
					},
					{
						Denom:    "uqck",
						Exponent: 6,
					},
				},
				Base:    s.defaultDenom,
				Display: "uqck",
				Name:    "OSMO",
				Symbol:  "OSMO",
			}),
			expectedPass: false,
		},
		{
			desc: "invalid metadata (missing display denom unit)",
			msgSetDenomMetadata: *types.NewMsgSetDenomMetadata(s.TestAccs[0].String(), banktypes.Metadata{
				Description: "yeehaw",
				DenomUnits: []*banktypes.DenomUnit{
					{
						Denom:    s.defaultDenom,
						Exponent: 0,
					},
				},
				Base:    s.defaultDenom,
				Display: "uqck",
				Name:    "OSMO",
				Symbol:  "OSMO",
			}),
			expectedPass: false,
		},
	} {
		s.Run(fmt.Sprintf("Case %s", tc.desc), func() {
			bankKeeper := s.App.BankKeeper
			res, err := s.msgServer.SetDenomMetadata(sdk.WrapSDKContext(s.Ctx), &tc.msgSetDenomMetadata)
			if tc.expectedPass {
				s.Require().NoError(err)
				s.Require().NotNil(res)

				md, found := bankKeeper.GetDenomMetaData(s.Ctx, s.defaultDenom)
				s.Require().True(found)
				s.Require().Equal(tc.msgSetDenomMetadata.Metadata.Name, md.Name)
			} else {
				s.Require().Error(err)
			}
		})
	}
}
