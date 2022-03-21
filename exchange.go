package exchange

type Tag string

const TagUSD = Tag("usd")

type Service interface {
	// Convert converts a given mount in source currency to a destination currency
	Convert(amount float64, source Tag, destination Tag) (float64, error)
}

type ExternalExchangeRatesService interface {
	// GetLiveRates returns live rates from third party sources
	GetLiveRates() map[Tag]float64
}
