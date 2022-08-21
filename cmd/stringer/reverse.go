package stringer

import (
	"fmt"

	"github.com/gain620/weatherctl/pkg/stringer"
	"github.com/spf13/cobra"
)

var ReverseCmd = &cobra.Command{
	Use:     "reverse",
	Aliases: []string{"rev"},
	Short:   "Reverses a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := stringer.Reverse(args[0])
		fmt.Println(res)
	},
}
