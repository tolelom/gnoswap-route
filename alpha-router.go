package main

import (
	"fmt"
	"gnoswap-route/core"
	"gnoswap-route/core/entities"
	"gnoswap-route/core/entities/fractions"
	"gnoswap-route/providers"
	v3 "gnoswap-route/providers/v3"
	"gnoswap-route/quoters"
)


type AlphaRouterParams struct {
	ChainId core.ChainId
	V3SubgraphProvider v3.IV3SubpGraphProvider
	V3PoolProvider v3.IV3PoolProvider
	OnChainQuoteProvider providers.IOnChainQuoteProvider
	TokenProvider providers.ITokenProvider
	PortionProvider providers.IPortionProvider
}

func mergeAlphaRouterParams(param ...AlphaRouterParams) AlphaRouterParams {
	var mergedParam AlphaRouterParams

	// 구현 대기



	return mergedParam
}


// TODO: 값이 어떤 범위 내에서 사용되는 지 확인 후 적절한 자료형으로 변환시켜야 할 수 있다.
type ProtocolPoolSelection struct {
	topN int64
	topNDirectSwaps int64
	topNTokenInOUt int64
	topNSecondHop int64
	//topNSecondHopForTokenAddress
	//tokensToAvoidOnSecondHops
	topNWithEachBaseToken int64
	topNWithBaseToken int64
}

type AlphaRouterConfig struct {
	v3PoolSelection ProtocolPoolSelection
	maxSwapsPerPath int64
	maxSplits int64
	minSplits int64
	distributionPercent int64
}


type AlphaRouter struct {
	chainId core.ChainId
	v3SubgraphProvider v3.IV3SubpGraphProvider
	v3PoolProvider v3.IV3PoolProvider
	onChainQuoteProvider providers.IOnChainQuoteProvider
	tokenProvider providers.ITokenProvider
	v3Quoter quoters.V3Quoter
	portionProvider providers.IPortionProvider
}

func NewAlphaRouter(alphaRouterParams AlphaRouterParams) *AlphaRouter {
	return &AlphaRouter{
		chainId: alphaRouterParams.ChainId,
		v3SubgraphProvider: alphaRouterParams.V3SubgraphProvider,
		v3PoolProvider: alphaRouterParams.V3PoolProvider,
		onChainQuoteProvider: alphaRouterParams.OnChainQuoteProvider,
		tokenProvider: alphaRouterParams.TokenProvider,
		portionProvider: alphaRouterParams.PortionProvider,

		v3Quoter: quoters.NewV3Quoter(
			alphaRouterParams.V3SubgraphProvider,
			alphaRouterParams.V3PoolProvider,
			alphaRouterParams.OnChainQuoteProvider,
			alphaRouterParams.TokenProvider,
			alphaRouterParams.ChainId),
	}
}

// 라우터의 핵심 함수로 볼 수 있고 내가 짠다면 run에 해당하는 함수인 만큼 내용을 함수화
// 시키면서 리팩토링을 해야할 것
// TODO: 원본 코드에서는 async 함수이다.
func (a AlphaRouter) route(
	amount fractions.CurrencyAmount,
	quoteCurrency entities.Currency, // 기준 통화
	tradeType core.TradeType, // TradeType은 EXACT_INPUT과 EXACT_OUTPUT 중 하나이다.
	) {


	originalAmount := amount // 토큰의 양


	// tradeType에 따라 amount.currency과 quoteCurrency를 적절히 넣어주는 기능
	var currencyIn, currencyOut entities.Currency
	// TODO: 일단 함수로 뺴지 않음, 추후에 다시 뺄 생각 해야함
	if (tradeType == core.EXACT_INPUT) {
		currencyIn = amount.Currency
		currencyOut = quoteCurrency
	}
	else {
		currencyIn = quoteCurrency
		currencyOut = amount.Currency
	}

	tokenIn := currencyIn
	tokenOut := currencyOut



	// TradeType은 EXACT_INPUT과 EXACT_OUTPUT 중 하나이다. 이게 무슨 역할인지는 고민해보자
	// EXACT_OUT은 사용자가 받고자 하는 토큰의 양을 지정하는 것이고
	// EXACT_IN은 사용자가 교환하고자 하는 토큰의 양을 지정하는 것이다.
	// 즉, 사용자가 받고 싶은 토큰의 양을 지정한 경우이고 이 때 amout값을 계산해서 넣어주는 거 같다.
	if (tradeType == core.EXACT_OUTPUT) {
		// portionAmount :=
		var portionAmount fractions.CurrencyAmount
		//if (portionAmount) {

		// 정확한 이해는 못했지만 원활한 거래를 위해 거래에서 사용되는 portion amount를 계산해서 더해주는 부분이다.
		// 여기에는 trading fees와 slippage 등이 포함 될 것으로 추측
		// amount += portionAmount
		// }
	}

	//routingConfig := mergeAlphaRouterParams()

	var quoteToken entities.Currency = quoteCurrency




	var tokenOutAmount fractions.CurrencyAmount
	if (tradeType == core.EXACT_OUTPUT) {
		tokenOutAmount = originalAmount
	}
	else {
		tokenOutAmount = quote
	}
}

// 라우팅 경로를 찾는 핵심적인 함수로 추측
func (a AlphaRouter) getSwapRouteFromChain(
	amount fractions.CurrencyAmount,
	tokenIn entities.Token,
	tokenOut entities.Token,
	quoteToken entities.Token,
	tradeType core.TradeType,
	routingConfig AlphaRouterConfig) {

	percents, amount := a.getAmountDistribution(amount, routingConfig)

	var v3CandidatePoolsPromise = getV3CandidatePools()
}



// 이게 왜 필요한걸까
func (a AlphaRouter) getAmountDistribution(amout fractions.CurrencyAmount, routingConfig AlphaRouterConfig) {

}


// TEST
func (a AlphaRouter) run() {
	fmt.Println("Alpha router run")
}

