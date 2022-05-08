package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var weatherCmd = &cobra.Command{
	Use:     "weather",
	Short:   "Get current weather info according to your current location",
	Example: "conctl weather [command] [options] [args]",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Running %v \n", args[0:])
		hostCmd := exec.Command(args[0], args[1:]...)
		hostCmd.Stdin = os.Stdin
		hostCmd.Stdout = os.Stdout
		hostCmd.Stderr = os.Stderr

		err := hostCmd.Run()
		if err != nil {
			log.Fatalf("Error running shell command : %v", err)
		}

	},
}
