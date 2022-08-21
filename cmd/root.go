/*
Copyright © 2022 Gain Chang gainchang620@gmail.com

*/
package cmd

import (
	"fmt"
	"github.com/gain620/weatherctl/cmd/stringer"
	"github.com/gain620/weatherctl/cmd/weather"
	"github.com/spf13/cobra"
	"log"
)

const version = "1.0.1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "weatherctl",
	Version: version,
	Short:   "A simple CLI program to get the current weather info",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: weatherctl [command] [flags]\n\nFor more information, use --help.")
	},
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get specific obejct",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("weatherctl can not execute rootCmd : %v", err)
	}
}

func init() {
	//rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(archCmd)
	//rootCmd.AddCommand(shellCmd)
	rootCmd.AddCommand(getCmd)

	rootCmd.AddCommand(stringer.ReverseCmd)
	rootCmd.AddCommand(stringer.InspectCmd)

	getCmd.AddCommand(weather.WeatherCmd)
	weather.WeatherCmd.Flags().StringP("city", "c", "seoul", "Input name of the city.")
	weather.WeatherCmd.Flags().StringP("aqi", "a", "yes", "Get air quality of the target location.")
	weather.WeatherCmd.Flags().StringP("temp", "t", "celsius", "Input type of the temperature. e.g,) weatherctl get weather --temp=celsius or -t fahrenheit")
}
