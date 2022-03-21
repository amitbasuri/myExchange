# My Exchange

My Exchange stores live exchange rates in local memory and sync
them with some external source via API e.g (Coingecko API or CurrencyLayer).

##### Build and run on local
```shell
go build -o ./bin/myExchange -i ./cmd/myExchange
./bin/myExchange
```
### Example Usage
```go
func main() {
    fiatRatesSvc := fiatExchange.NewFiatExternalExchangeRatesService("__currency_layer_api_token__")
    cryptoRatesSvc := cryptoExchange.NewCryptoExternalExchangeRatesService()

    myExchange := storage.Exchange(fiatRatesSvc, cryptoRatesSvc)
    convertedAmt, _ := myExchange.Convert(1.54,"btc", "eur")
    //....
}
```

#### Sample Run
```shell
100 eur is equal to 0.00000726 ltc
200 ltc is equal to 2311813710.73422527 gbp
300 eur is equal to 0.0005528 eth
400 eur is equal to 595.98827369 aud
500 usd is equal to 673.4295 aud
600 eth is equal to 15227.01917088 ltc
700 inr is equal to 8.3317039 eur
800 usd is equal to 60757.32 inr
900 gbp is equal to 0.02855099 btc
```
