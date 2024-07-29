package quoters

import (
	"gnoswap-route/core"
	"gnoswap-route/providers"
	v3 "gnoswap-route/providers/v3"
)

type V3Quoter struct {
	v3SubgraphPrivvider  v3.IV3SubpGraphProvider
	v3PoolProvider       v3.IV3PoolProvider
	onChainQuoteProvider providers.IOnChainQuoteProvider
}

func NewV3Quoter(
	v3SubgraphProvider v3.IV3SubpGraphProvider,
	v3PoolProvider v3.IV3PoolProvider,
	onChainQuoteProvider providers.IOnChainQuoteProvider,
	tokenProvider providers.ITokenProvider, // TODO: BaseQuoter에서 사용
	chainId core.ChainId, // TODO: BaseQuoter에서 사용
) V3Quoter {
	return V3Quoter{
		v3SubgraphPrivvider:  v3SubgraphProvider,
		v3PoolProvider:       v3PoolProvider,
		onChainQuoteProvider: onChainQuoteProvider,
	}
}
