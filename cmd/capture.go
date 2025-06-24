package cmd

import (
	"fmt"

	"github.com/bhh32/go-shark/pkg/capture"
	"github.com/spf13/cobra"
)

var iface string
var duration int
var filter string

var CaptureCmd = &cobra.Command{
	Use:   "capture",
	Short: "Capture packets on a network interface",
	Run: func(cmd *cobra.Command, args []string) {
		if iface == "" {
			fmt.Println("Please specify an interface with --iface")
			return
		}

		if duration <= 0 {
			fmt.Println("Duration cannot be less than or equal to 0\nUsing default duration of 10 seconds")
			duration = 10
		}

		capture.CapturePackets(iface, duration, filter)
	},
}
