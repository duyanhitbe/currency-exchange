package config

import "fmt"

const (
	BaseURL = "https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1"
)

func GetExchangeEndpoint(currency string) string {
	return fmt.Sprintf("%s/currencies/%s.json", BaseURL, currency)
}

func GetListCurrencyEndpoint() string {
	return fmt.Sprintf("%s/currencies.json", BaseURL)
}
