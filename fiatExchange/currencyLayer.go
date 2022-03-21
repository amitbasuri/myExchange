package fiatExchange

import (
	"encoding/json"
	"exchange"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type currencyLayerExternalExchangeRatesServiceImp struct {
	accessKey string
}

type currencyLayerResponse struct {
	Success bool               `json:"success"`
	Quotes  map[string]float64 `json:"quotes"`
}

// GetLiveRates fetches live rates from currencyLayer
func (cg *currencyLayerExternalExchangeRatesServiceImp) GetLiveRates() map[exchange.Tag]float64 {
	path := fmt.Sprintf("http://api.currencylayer.com/live?access_key=%s&format=1", cg.accessKey)
	response, err := http.Get(path)
	if err != nil {
		log.Printf("error calling currency layer api %s \n", err)
		return nil
	}

	if response.StatusCode != http.StatusOK {
		log.Printf("error calling currency layer api response status %d \n", response.StatusCode)
		return nil
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("error GetLiveRates %s \n", err)
		return nil
	}

	var responseObject currencyLayerResponse
	json.Unmarshal(responseData, &responseObject)

	exchangeRates := make(map[exchange.Tag]float64)
	for qoute, rate := range responseObject.Quotes {
		convertTo := exchange.Tag(strings.ToLower(qoute[3:]))
		exchangeRates[convertTo] = rate
	}

	return exchangeRates
}

func NewFiatExternalExchangeRatesService(accessKey string) exchange.ExternalExchangeRatesService {
	return &currencyLayerExternalExchangeRatesServiceImp{
		accessKey: accessKey,
	}
}
