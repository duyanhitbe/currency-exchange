/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strings"

	"github.com/duyanhitbe/currency-exchange/internal/api"
	"github.com/duyanhitbe/currency-exchange/internal/config"
	"github.com/duyanhitbe/currency-exchange/internal/helpers"
	"github.com/spf13/cobra"
)

// ratesCmd represents the rates command
var ratesCmd = &cobra.Command{
	Use:   "rates [code]",
	Short: "List all currency rates compare with [code]",
	Long: `
Arguments:
  code  Currency code of country you want to compare.`,
	Run:  runRates,
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(ratesCmd)
	ratesCmd.Flags().StringVarP(&currenciesFlag, "currencies", "c", "", "Comma-separated list of currency codes to filter")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ratesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ratesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runRates(cmd *cobra.Command, args []string) {
	code := args[0]

	// Get the exchange rate endpoint
	url := config.GetExchangeEndpoint(code)
	data := api.GET(url)

	// Assert the type of data[code] as map[string]interface{}
	currencyData, ok := data[code].(map[string]interface{})
	if !ok {
		log.Fatal("Cannot parse data for the provided currency code")
	}

	// Split the currenciesFlag by comma if it is set
	filteredCurrencies := map[string]bool{}
	if currenciesFlag != "" {
		for _, c := range strings.Split(currenciesFlag, ",") {
			filteredCurrencies[strings.TrimSpace(c)] = true
		}
	}

	w := helpers.NewTabWriter()
	defer w.Writer.Flush()

	// Print out the map
	w.WriteHeaderListRates()
	i := 0
	// Iterate over the currency data and print only the filtered rates
	for c, r := range currencyData {
		// Check if we are filtering currencies, and if this currency is in the filter
		if len(filteredCurrencies) > 0 && !filteredCurrencies[c] {
			continue
		}

		// Try to convert the rate to float64, as the rates are in float64 format
		rate, ok := r.(float64)
		if !ok {
			log.Printf("Cannot parse rate for currency: %s", c)
			continue
		}
		w.WriteListRates(i, c, rate)

		i++
	}
}
