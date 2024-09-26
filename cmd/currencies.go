/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/duyanhitbe/currency-exchange/internal/api"
	"github.com/duyanhitbe/currency-exchange/internal/config"
	"github.com/duyanhitbe/currency-exchange/internal/helpers"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "currencies",
	Short: "List all currencies",
	Run:   runCurrencies,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runCurrencies(cmd *cobra.Command, args []string) {
	// API endpoint
	url := config.GetListCurrencyEndpoint()

	result := api.GET(url)

	var rows [][]string

	for code, country := range result {
		rows = append(rows, []string{fmt.Sprintf("%s", country), code})
	}

	table := helpers.NewTable()
	table.WriteCurrencies(rows)
	table.Print()
}
