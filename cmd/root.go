/*
Copyright © 2022 Gain Chang gainchang620@gmail.com

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "conctl",
	Short: "A sample CLI program to demonstrate how container really works!",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Usage: conctl [command] [flags]\n\nFor more information, use --help.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("conctl can not execute rootCmd : %v", err)
	}
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(archCmd)
	rootCmd.AddCommand(shellCmd)
}
