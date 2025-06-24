package cmd

import (
	"fmt"
	"os"

	"github.com/bhh32/go-shark/pkg/tracer"
	"github.com/spf13/cobra"
)

var maxHops int
var dest string

var TracerCmd = &cobra.Command{
	Use:   "tracer",
	Short: "Perform a trace to a host",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dest := args[0]
		if err := tracer.RunTrace(dest, maxHops); err != nil {
			fmt.Fprintln(os.Stderr, "Tracer error:", err)
		}
	},
}
