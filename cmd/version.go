package cmd

import (
	"fmt"
	"github.com/gain620/container-deepdive/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Output the version of CLI program.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("%s version is %s", config.AppName, config.Version))
	},
}
