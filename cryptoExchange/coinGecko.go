package cryptoExchange

import (
	"encoding/json"
	"exchange"
	"io/ioutil"
	"log"
	"net/http"
)

type coinGeckoExternalExchangeRatesServiceImp struct{}

type coinGeckoExchangeResponse struct {
	Rates map[string]coinGeckoExchange `json:"rates"`
}

type coinGeckoExchange struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

// GetLiveRates fetches live rates from coingecko
func (cg *coinGeckoExternalExchangeRatesServiceImp) GetLiveRates() map[exchange.Tag]float64 {
	path := "https://api.coingecko.com/api/v3/exchange_rates"
	response, err := http.Get(path)
	if err != nil {
		log.Printf("error calling coin gecko api %s \n", err)
		return nil
	}

	if response.StatusCode != http.StatusOK {
		log.Printf("error calling coin gecko api response status %d \n", response.StatusCode)
		return nil
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("error GetLiveRates %s \n", err)
		return nil
	}

	var responseObject coinGeckoExchangeResponse
	json.Unmarshal(responseData, &responseObject)

	exchangeRates := make(map[exchange.Tag]float64)
	for convertTo, rate := range responseObject.Rates {
		if rate.Type == "crypto" {
			exchangeRates[exchange.Tag(convertTo)] = rate.Value

		}
		if convertTo == string(exchange.TagUSD) {
			exchangeRates[exchange.Tag(convertTo)] = rate.Value
		}
	}

	return exchangeRates

}

func NewCryptoExternalExchangeRatesService() exchange.ExternalExchangeRatesService {

	return &coinGeckoExternalExchangeRatesServiceImp{}
}
