package fractions

import "gnoswap-route/core/entities"

type CurrencyAmount struct {
	Currency entities.Currency
}

func (a CurrencyAmount) Add(other CurrencyAmount) CurrencyAmount {
	return CurrencyAmount{}
}
