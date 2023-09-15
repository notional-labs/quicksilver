package stableswap

import (
	"errors"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/quicksilver-zone/quicksilver/third-party-chains/osmosis-types/gamm"
	"github.com/quicksilver-zone/quicksilver/third-party-chains/osmosis-types/gamm/pool-models/internal/cfmm_common"
)

var (
	cubeRootTwo, _   = sdk.NewDec(2).ApproxRoot(3)
	threeCubeRootTwo = cubeRootTwo.MulInt64(3)
)

// solidly CFMM is xy(x^2 + y^2) = k
// So we want to solve for a given addition of `b` units of y into the pool,
// how many units `a` of x do we get out.
// So we solve the following expression for `a`
// xy(x^2 + y^2) = (x - a)(y + b)((x - a)^2 + (y + b)^2)
func solveCfmm(xReserve, yReserve, yIn sdk.Dec) sdk.Dec {
	if !yReserve.Add(yIn).IsPositive() {
		panic("invalid yReserve, yIn combo")
	}

	// use the following wolfram alpha link to solve the equation
	// https://www.wolframalpha.com/input?i=solve+for+a%2C+xy%28x%5E2+%2B+y%5E2%29+%3D+%28x+-+a%29%28y+%2B+b%29%28%28x+-+a%29%5E2+%2B+%28y+%2Bb%29%5E2%29+
	// This returns (copied from wolfram):
	// assuming (correctly) that b + y!=0
	// a = (-27 b^2 x^3 y - 27 b^2 x y^3 + sqrt((-27 b^2 x^3 y - 27 b^2 x y^3 - 54 b x^3 y^2 - 54 b x y^4 - 27 x^3 y^3 - 27 x y^5)^2 + 4 (3 b^4 + 12 b^3 y + 18 b^2 y^2 + 12 b y^3 + 3 y^4)^3) - 54 b x^3 y^2 - 54 b x y^4 - 27 x^3 y^3 - 27 x y^5)^(1/3)/(3 2^(1/3) (b + y)) - (2^(1/3) (3 b^4 + 12 b^3 y + 18 b^2 y^2 + 12 b y^3 + 3 y^4))/(3 (b + y) (-27 b^2 x^3 y - 27 b^2 x y^3 + sqrt((-27 b^2 x^3 y - 27 b^2 x y^3 - 54 b x^3 y^2 - 54 b x y^4 - 27 x^3 y^3 - 27 x y^5)^2 + 4 (3 b^4 + 12 b^3 y + 18 b^2 y^2 + 12 b y^3 + 3 y^4)^3) - 54 b x^3 y^2 - 54 b x y^4 - 27 x^3 y^3 - 27 x y^5)^(1/3)) + (b x + x y)/(b + y) and b + y!=0
	// We simplify and separate out terms to get that its the following:
	// The key substitutions are that 3(b+y)^4 = 3 b^4 + 12 b^3 y + 18 b^2 y^2 + 12 b y^3 + 3 y^4
	// and -27 x y (b + y)^2 (x^2 + y^2) = -27 b^2 x^3 y - 27 b^2 x y^3 - 54 b x^3 y^2 - 54 b x y^4 - 27 x^3 y^3 - 27 x y^5
	// I added {} myself, making better distinctions between entirely distinct terms.
	// a = {(-27 x y (b + y)^2 (x^2 + y^2)
	//			+ sqrt(
	//				(-27 x y (b + y)^2 (x^2 + y^2))^2
	//				+ 108 ((b+y)^4)^3
	// 			)^(1/3)
	// 		  / (3 2^(1/3) (b + y))}
	//		- {(2^(1/3) (3 (b + y)^4))
	//		  /(3 (b + y)
	// 			(-27 x y (b + y)^2 (x^2 + y^2)
	// 				+ sqrt(
	//					(-27 x y (b + y)^2 (x^2 + y^2))^2
	//					+ 108 ((b+y)^4)^3)
	// 			)^(1/3))}
	//      + {(b x + x y)/(b + y)}
	// we further simplify, and call:
	// foo = (-27 x y (b + y)^2 (x^2 + y^2)
	// 			+ sqrt(
	//				(-27 x y (b + y)^2 (x^2 + y^2))^2
	//				+ 108 ((b+y)^4)^3)
	//		 )^(1/3)
	// Thus, a is then:
	// a = {foo / (3 2^(1/3) (b + y))}
	//		- {(3 * 2^(1/3) (b+y)^4)
	//		  /(3 (b + y) foo)}
	//      + {(b x + x y)/(b + y)}
	// Let:
	// term1 = {foo / (3 2^(1/3) (b + y))}
	// term2 = {(3 * 2^(1/3) (b+y)^4) /(3 (b + y) foo)} =  2^(1/3) (b+y)^3 / foo
	// term3 = {(b x + x y)/(b + y)}

	// prelude, compute all the xy cross terms. Consider keeping these precomputed in the struct,
	// and maybe in state.
	x := xReserve
	y := yReserve
	x2py2 := x.Mul(x).AddMut(y.Mul(y))

	xy := x.Mul(y)

	b := yIn

	bpy := b.Add(y)
	bpy2 := bpy.Mul(bpy)
	bpy3 := bpy2.Mul(bpy)
	bpy4 := bpy2.Mul(bpy2)

	// TODO: Come back and optimize alot of the calculations

	// Now we compute foo
	// foo = (-27 x y (b + y)^2 (x^2 + y^2)
	// 			+ sqrt(
	//				(-27 x y (b + y)^2 (x^2 + y^2))^2
	//				+ 108 ((b+y)^4)^3)
	//		 )^(1/3)
	// This has a y^12 term in it, which is unappealing, so we spend some energy reducing this max bitlen.
	// foo = (-27 x y (b + y)^2 (x^2 + y^2)
	// 			+ (b + y)^2 sqrt(
	//				729 (x y (x^2 + y^2))^2
	//				+ 108 (b+y)^8)
	//		 )^(1/3)
	// let e = x y (x^2 + y^2))
	// foo = (-27 (b + y)^2 e
	// 			+ (b + y)^2 sqrt(
	//				729 e^2 + 108 (b+y)^8)
	//		 )^(1/3)

	e := xy.Mul(x2py2) // xy(x^2 + y^2)

	// t1 = -27 (b + y)^2 e
	t1 := e.Mul(bpy2).MulInt64Mut(-27)

	// compute d = (b + y)^2 sqrt(729 e^2 + 108 (b+y)^8)
	bpy8 := bpy4.Mul(bpy4)
	sqrt_inner := e.MulMut(e).MulInt64Mut(729).AddMut(bpy8.MulInt64Mut(108)) // 729 e^2 + 108 (b+y)^8
	sqrt, err := sqrt_inner.ApproxSqrt()
	if err != nil {
		panic(err)
	}
	d := sqrt.MulMut(bpy2)

	// foo = (t1 + d)^(1/3)
	foo3 := t1.AddMut(d)
	foo, _ := foo3.ApproxRoot(3)

	// a = {foo / (3 2^(1/3) (b + y))}
	//		- {(2^(1/3) banana) / (3 (b + y) foo}
	//      + {(b x + x y)/(b + y)}

	// term1 := {foo / (3 2^(1/3) (b + y))}
	term1Denominator := threeCubeRootTwo.Mul(bpy)
	term1 := foo.Quo(term1Denominator)
	// term2 := {(2^(1/3) (b+y)^3) / (foo}
	term2 := cubeRootTwo.Mul(bpy3)
	term2 = term2.Quo(foo)
	// term3 := {(b x + x y)/(b + y)}
	term3Numerator := b.Mul(x).Add(xy)
	term3 := term3Numerator.Quo(bpy)

	a := term1.Sub(term2).Add(term3)
	return a
}

func spotPrice(baseReserve, quoteReserve sdk.Dec) sdk.Dec {
	// y = baseAsset, x = quoteAsset
	// Define f_{y -> x}(a) as the function that outputs the amount of tokens X you'd get by
	// trading "a" units of Y against the pool, assuming 0 swap fee, at the current liquidity.
	// The spot price of the pool is then lim a -> 0, f_{y -> x}(a) / a
	// For uniswap f_{y -> x}(a) = x - xy/(y + a),
	// The spot price equation of y in terms of x is X_SUPPLY/Y_SUPPLY.
	// You can work out that it follows from the above relation!
	//
	// Now we have to work this out for the much more complex CFMM xy(x^2 + y^2).
	// Or we can sidestep this, by just picking a small value a, and computing f_{y -> x}(a) / a,
	// and accept the precision error.

	// We arbitrarily choose a = 1, and anticipate that this is a small value at the scale of
	// xReserve & yReserve.
	a := sdk.OneDec()
	// no need to divide by a, since a = 1.
	return solveCfmm(baseReserve, quoteReserve, a)
}

// returns outAmt as a decimal
func (p *Pool) calcOutAmtGivenIn(tokenIn sdk.Coin, tokenOutDenom string, swapFee sdk.Dec) (sdk.Dec, error) {
	reserves, err := p.getScaledPoolAmts(tokenIn.Denom, tokenOutDenom)
	if err != nil {
		return sdk.Dec{}, err
	}
	tokenInSupply, tokenOutSupply := reserves[0], reserves[1]
	// We are solving for the amount of token out, hence x = tokenOutSupply, y = tokenInSupply
	cfmmOut := solveCfmm(tokenOutSupply, tokenInSupply, sdk.NewDecFromInt(tokenIn.Amount))
	outAmt := p.getDescaledPoolAmt(tokenOutDenom, cfmmOut)
	return outAmt, nil
}

// returns inAmt as a decimal
func (p *Pool) calcInAmtGivenOut(tokenOut sdk.Coin, tokenInDenom string, swapFee sdk.Dec) (sdk.Dec, error) {
	reserves, err := p.getScaledPoolAmts(tokenInDenom, tokenOut.Denom)
	if err != nil {
		return sdk.Dec{}, err
	}
	tokenInSupply, tokenOutSupply := reserves[0], reserves[1]
	// We are solving for the amount of token in, cfmm(x,y) = cfmm(x + x_in, y - y_out)
	// x = tokenInSupply, y = tokenOutSupply, yIn = -tokenOutAmount
	cfmmIn := solveCfmm(tokenInSupply, tokenOutSupply, sdk.NewDecFromInt(tokenOut.Amount).Neg())
	inAmt := p.getDescaledPoolAmt(tokenInDenom, cfmmIn.NegMut())
	return inAmt, nil
}

func (p *Pool) calcSingleAssetJoinShares(tokenIn sdk.Coin, swapFee sdk.Dec) (sdkmath.Int, error) {
	poolWithAddedLiquidityAndShares := func(newLiquidity sdk.Coin, newShares sdkmath.Int) gamm.PoolI {
		paCopy := p.Copy()
		paCopy.updatePoolForJoin(sdk.NewCoins(tokenIn), newShares)
		return &paCopy
	}
	// TODO: Correctly handle swap fee
	return cfmm_common.BinarySearchSingleAssetJoin(p, tokenIn, poolWithAddedLiquidityAndShares)
}

// We can mutate pa here
// TODO: some day switch this to a COW wrapped pa, for better perf
func (p *Pool) joinPoolSharesInternal(ctx sdk.Context, tokensIn sdk.Coins, swapFee sdk.Dec) (numShares sdkmath.Int, newLiquidity sdk.Coins, err error) {
	if len(tokensIn) == 1 {
		numShares, err = p.calcSingleAssetJoinShares(tokensIn[0], swapFee)
		newLiquidity = tokensIn
		return numShares, newLiquidity, err
	} else if len(tokensIn) != p.NumAssets() || !tokensIn.DenomsSubsetOf(p.GetTotalPoolLiquidity(ctx)) {
		return sdk.ZeroInt(), sdk.NewCoins(), errors.New(
			"stableswap pool only supports LP'ing with one asset, or all assets in pool")
	}

	// Add all exact coins we can (no swap). ctx arg doesn't matter for Stableswap
	numShares, remCoins, err := cfmm_common.MaximalExactRatioJoin(p, sdk.Context{}, tokensIn)
	if err != nil {
		return sdk.ZeroInt(), sdk.NewCoins(), err
	}
	p.updatePoolForJoin(tokensIn.Sub(remCoins...), numShares)

	for _, coin := range remCoins {
		// TODO: Perhaps add a method to skip if this is too small.
		newShare, err := p.calcSingleAssetJoinShares(coin, swapFee)
		if err != nil {
			return sdk.ZeroInt(), sdk.NewCoins(), err
		}
		p.updatePoolForJoin(sdk.NewCoins(coin), newShare)
		numShares = numShares.Add(newShare)
	}

	return numShares, tokensIn, nil
}
