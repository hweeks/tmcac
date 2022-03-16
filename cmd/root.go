/*
Copyright © 2022 hweeks tmcac_root@hweeks.com

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tmcac",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func handle_arg_input(cmd *cobra.Command, args []string) {
	var price_range string
	var market_location string
	var vehicle_type string
	for i, str := range args {
		switch i {
		case 0:
			price_range = str
		case 1:
			market_location = str
		case 2:
			vehicle_type = str
		}
	}
	if price_range != "" {
		fmt.Println(price_range)
	}
	if market_location != "" {
		fmt.Println(market_location)
	}
	if vehicle_type != "" {
		fmt.Println(vehicle_type)
	}
}

func init() {
	var cmd_price_range = &cobra.Command{
		Use:   "range [cost, location, vehicle-type]",
		Short: "Find a list of profitable vehicles to own in an area",
		Long:  `range is a tool designed to help myself find well priced vehicles`,
		Args:  cobra.MinimumNArgs(3),
		Run:   handle_arg_input,
	}
	rootCmd.AddCommand(cmd_price_range)
}
