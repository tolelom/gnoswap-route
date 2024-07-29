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


// TODO: 원본 코드에서는 async 함수이다.
func (a AlphaRouter) route(
	amount fractions.CurrencyAmount,
	quoteCurrency entities.Currency,
	tradeType core.TradeType) {


	originalAmount := amount

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

	if (tradeType == core.EXACT_OUTPUT) {
		// portionAmount :=
		var portionAmount fractions.CurrencyAmount
		//if (portionAmount)
		//amount += portionAmount
	}

	//routingConfig := mergeAlphaRouterParams()





	var tokenOutAmount fractions.CurrencyAmount
	if (tradeType == core.EXACT_OUTPUT) {
		tokenOutAmount = originalAmount
	}
	else {
		tokenOutAmount = quote
	}
}



// TEST
func (a AlphaRouter) run() {
	fmt.Println("Alpha router run")
}

