package storage

import (
	"exchange"
	"sync"
)

const cryptoToFiatRatesBaseTag = exchange.TagUSD

type exchangeImpl struct {
	fiatRates                   sync.Map
	cryptoRates                 sync.Map
	fiatRatesService            exchange.ExternalExchangeRatesService
	cryptoRatesService          exchange.ExternalExchangeRatesService
	ratesRefreshIntervalSeconds int
}

func (ex *exchangeImpl) isCrypto(tag exchange.Tag) bool {
	if tag == exchange.TagUSD {
		return false
	}
	_, ok := ex.cryptoRates.Load(tag)
	return ok
}

func (ex *exchangeImpl) isFiat(tag exchange.Tag) bool {
	_, ok := ex.fiatRates.Load(tag)
	return ok
}

func (ex *exchangeImpl) getCryptoRate(source exchange.Tag) float64 {
	v, _ := ex.cryptoRates.Load(source)
	return v.(float64)
}

func (ex *exchangeImpl) getFiatRate(source exchange.Tag) float64 {
	v, _ := ex.fiatRates.Load(source)
	return v.(float64)
}

func (ex *exchangeImpl) storeCryptoExchangeRate(tag exchange.Tag, rate float64) {
	ex.cryptoRates.Store(tag, rate)
}

func (ex *exchangeImpl) storeFiatExchangeRate(tag exchange.Tag, rate float64) {
	ex.fiatRates.Store(tag, rate)
}
