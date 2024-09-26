/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/duyanhitbe/currency-exchange/internal/api"
	"github.com/duyanhitbe/currency-exchange/internal/config"
	"github.com/duyanhitbe/currency-exchange/internal/helpers"
	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert [amount] [from] [to]",
	Short: "Convert your currency to provided country",
	Long: `
Arguments:
  amount  Amount of money you want to convert.
  from  Currency code of amount.
  amount  Currency code of country you want to convert.`,
	Run:  runConvert,
	Args: cobra.ExactArgs(3),
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runConvert(cmd *cobra.Command, args []string) {
	a := args[0]
	from := args[1]
	t := args[2]

	listTo := strings.Split(t, ",")

	amount, err := strconv.ParseFloat(a, 64)
	if err != nil {
		log.Fatalf("Invalid amount %v", err)
	}

	// Get the exchange rate endpoint
	url := config.GetExchangeEndpoint(from)
	data := api.GET(url)

	// Assert the type of data[code] as map[string]interface{}
	currencyData, ok := data[from].(map[string]interface{})
	if !ok {
		log.Fatal("Cannot parse data for the provided currency code")
	}

	var rows [][]string
	for _, to := range listTo {
		r := currencyData[to]
		rate, ok := r.(float64)
		if !ok {
			log.Println("Cannot parse rate for currency")
		}
		rows = append(rows, []string{from, to, fmt.Sprintf("%f", amount), fmt.Sprintf("%f", amount*rate)})
	}

	table := helpers.NewTable()
	table.WriteConvert(rows)
	table.Print()
}
