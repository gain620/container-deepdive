package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"syscall"
)

// https://www.youtube.com/watch?v=Utf-A4rODH8
var shellCmd = &cobra.Command{
	Use:     "shell",
	Short:   "Use host's bash command line",
	Example: "conctl shell [command] [options] [args]",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		hostCmd := exec.Command("/proc/self/exe", append([]string{"childProc"}, args[1:]...)...)
		hostCmd.Stdin = os.Stdin
		hostCmd.Stdout = os.Stdout
		hostCmd.Stderr = os.Stderr

		// create namespace
		hostCmd.SysProcAttr = &syscall.SysProcAttr{
			Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
		}

		err := hostCmd.Run()
		if err != nil {
			log.Fatalf("Error running shell command : %v", err)
		}

	},
}

func chileProc() {
	fmt.Printf("Running %v as PID %d \n", os.Args[2:], os.Getpid())
	hostCmd := exec.Command(os.Args[2], os.Args[3:]...)
	hostCmd.Stdin = os.Stdin
	hostCmd.Stdout = os.Stdout
	hostCmd.Stderr = os.Stderr

	syscall.Chroot("/home/ubuntu")
	err := os.Chdir("/")
	if err != nil {
		log.Fatalf("Error running shell command : %v", err)
	}
	syscall.Mount("proc", "proc", "proc", 0, "")

	err = hostCmd.Run()
	if err != nil {
		log.Fatalf("Error running shell command : %v", err)
	}
}
