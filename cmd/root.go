package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "go-shark",
	Short: "A modern, fast packet analyzer",
	Long:  `Go-Shark is a WireShark-inspired packet analyzer written in Go.`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	captureCmd := CaptureCmd
	captureCmd.Flags().StringVar(&iface, "iface", "", "Network interface to capture from")
	captureCmd.Flags().IntVar(&duration, "duration", 10, "Capture duration in seconds")
	captureCmd.Flags().StringVar(&filter, "filter", "", "Filter what is captured")
	RootCmd.AddCommand(captureCmd)

	listIface := ListCmd
	RootCmd.AddCommand(listIface)

	tracerCmd := TracerCmd
	tracerCmd.Flags().StringVar(&dest, "dest", "", "Destination of the trace route")
	tracerCmd.Flags().IntVar(&maxHops, "max-hops", 30, "Maximum amout a hops a trace should take")
	RootCmd.AddCommand(tracerCmd)
}
