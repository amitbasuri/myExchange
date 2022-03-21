package storage

import (
	"errors"
	"exchange"
	"log"
	"sync"
	"time"
)

// exchangeImplInstance is a Singleton implementation
var exchangeImplInstance *exchangeImpl = nil

// Exchange returns new exchange Service
func Exchange(fiatRatesService, cryptoRatesService exchange.ExternalExchangeRatesService) exchange.Service {
	if exchangeImplInstance != nil {
		return exchangeImplInstance
	}
	exchangeImplInstance = &exchangeImpl{
		fiatRates:                   sync.Map{},
		cryptoRates:                 sync.Map{},
		fiatRatesService:            fiatRatesService,
		cryptoRatesService:          cryptoRatesService,
		ratesRefreshIntervalSeconds: 10,
	}
	exchangeImplInstance.setRates()

	// update rates in background
	go exchangeImplInstance.updateRatesAfterInterval()

	return exchangeImplInstance
}

// Convert converts a given mount in source currency to a destination currency
func (ex *exchangeImpl) Convert(amount float64, from exchange.Tag, convertTo exchange.Tag) (float64, error) {
	fromIsCrypto := ex.isCrypto(from)
	fromIsFiat := ex.isFiat(from)

	if fromIsCrypto == false && fromIsFiat == false {
		// if from tag is neither crypto nor fiat return error
		return 0, errors.New("invalid convert from")
	}
	convertToIsCrypto := ex.isCrypto(convertTo)
	convertToIsFiat := ex.isFiat(convertTo)
	if convertToIsCrypto == false && convertToIsFiat == false {
		// if destination tag is neither crypto nor fiat return error
		return 0, errors.New("invalid convert to")
	}
	rate := float64(0)

	if fromIsCrypto && convertToIsCrypto {
		rate = ex.getCryptoRate(convertTo) / ex.getCryptoRate(from)
	}
	if fromIsFiat && convertToIsFiat {
		rate = ex.getFiatRate(convertTo) / ex.getFiatRate(from)
	}
	if fromIsCrypto && convertToIsFiat {
		rate = ex.getCryptoRate(from) * ex.getCryptoRate(cryptoToFiatRatesBaseTag) * ex.getFiatRate(convertTo)
	}
	if fromIsFiat && convertToIsCrypto {
		rate = 1 / (ex.getFiatRate(from) * ex.getCryptoRate(cryptoToFiatRatesBaseTag) * ex.getCryptoRate(convertTo))
	}

	return amount * rate, nil
}

// setRates updates the rates in exchangeImplInstance
func (ex *exchangeImpl) setRates() {
	fiatRates := ex.fiatRatesService.GetLiveRates()
	cryptoRates := ex.cryptoRatesService.GetLiveRates()
	if fiatRates != nil {
		for curr, rate := range fiatRates {
			ex.storeFiatExchangeRate(curr, rate)
		}
	} else {
		log.Println("fiat Rates are not latest")
	}

	if cryptoRates != nil {
		for crypto, rate := range cryptoRates {
			ex.storeCryptoExchangeRate(crypto, rate)
		}
	} else {
		log.Println("crypto Rates are not latest")
	}

}

// updateRatesAfterInterval runs setRates after every given time interval
func (ex *exchangeImpl) updateRatesAfterInterval() {
	duration := time.Second * time.Duration(ex.ratesRefreshIntervalSeconds)
	for range time.Tick(duration) {
		ex.setRates()
	}
}
