package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-shark",
	Short: "A modern, fast packet analyzer",
	Long:  `Go-Shark is a WireShark-inspired packet analyzer written in Go.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
