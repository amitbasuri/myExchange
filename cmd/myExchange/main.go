package main

import (
	"exchange"
	"exchange/cryptoExchange"
	"exchange/fiatExchange"
	"exchange/storage"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fiatRatesSvc := fiatExchange.NewFiatExternalExchangeRatesService("dba1506d18c0e896e244bf4fec98b82f")
	cryptoRatesSvc := cryptoExchange.NewCryptoExternalExchangeRatesService()

	myExchange := storage.Exchange(fiatRatesSvc, cryptoRatesSvc)

	currencies := [8]string{"usd", "eur", "inr", "gbp", "aud", "btc", "eth", "ltc"}

	//loop over random currencies and convert from one to another for a given amount
	for i := 1; i < 10; i++ {
		time.Sleep(time.Second * 2)
		from := rand.Intn(len(currencies))
		to := rand.Intn(len(currencies))
		amount := i * 100
		convertAmount, err := myExchange.Convert(float64(amount), exchange.Tag(currencies[from]), exchange.Tag(currencies[to]))
		if err != nil {
			fmt.Printf("Error converting %d %s to %s, err: %v\n", amount, currencies[from], currencies[to], err)
		}
		fmt.Printf("%d %s is equal to %s %s\n", amount, currencies[from], floatFormat(convertAmount), currencies[to])
	}

}

func floatFormat(num float64) string {
	s := fmt.Sprintf("%.8f", num)
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}
