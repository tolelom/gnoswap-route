package functions

import (
	"gnoswap-route/core"
	"gnoswap-route/core/entities"
	"gnoswap-route/providers"
	v3 "gnoswap-route/providers/v3"
)

type V3GetCandidatePoolsParams struct {
	tokenIn   entities.Token
	tokenOut  entities.Token
	routeType core.TradeType
	//routingConfig AlphaRouterConfig
	subgraphProvider v3.IV3SubpGraphProvider
	tokenProvider    providers.ITokenProvider
	poolProvider     v3.IV3PoolProvider
	chainId          core.ChainId
}

func getV3CandidiatePools(v3GetCandidatePoolsParams V3GetCandidatePoolsParams) {

}
