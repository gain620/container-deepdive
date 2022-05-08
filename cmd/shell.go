package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"syscall"
)

var shellCmd = &cobra.Command{
	Use:     "shell",
	Short:   "Use host's bash command line",
	Example: "conctl shell [command] [options] [args]",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Running %v \n", args[0:])
		hostCmd := exec.Command(args[0], args[1:]...)
		hostCmd.Stdin = os.Stdin
		hostCmd.Stdout = os.Stdout
		hostCmd.Stderr = os.Stderr
		hostCmd.SysProcAttr = &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWUTS,
		}

		err := hostCmd.Run()
		if err != nil {
			log.Fatalf("Error running shell command : %v", err)
		}

	},
}
