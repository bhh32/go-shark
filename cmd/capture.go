package cmd

import (
	"fmt"

	"github.com/bhh32/go-shark/pkg/capture"
	"github.com/spf13/cobra"
)

var iface string
var duration int

var captureCmd = &cobra.Command{
	Use:   "capture",
	Short: "Capture packets on a network interface",
	Run: func(cmd *cobra.Command, args []string) {
		if iface == "" {
			fmt.Println("Please specify an interface with --iface")
			return
		}

		if duration <= 0 {
			duration = 10
		}

		capture.CapturePackets(iface, duration)
	},
}

func init() {
	captureCmd.Flags().StringVar(&iface, "iface", "", "Network interface to capture from")
	captureCmd.Flags().IntVar(&duration, "duration", 10, "Capture duration in seconds")
	rootCmd.AddCommand(captureCmd)
}
