package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var archCmd = &cobra.Command{
	Use:   "host",
	Short: "Output the general info about current host, such as architecture, OS, etc.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("OS: %s \n", runtime.GOOS)
		fmt.Printf("Arch: %s \n", runtime.GOARCH)
		fmt.Printf("CPU: %v \n", runtime.NumCPU())
		fmt.Printf("go version: %s \n", runtime.Version())
	},
}
