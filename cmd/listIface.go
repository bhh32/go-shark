package cmd

import (
	"fmt"

	"github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "ifaces",
	Short: "List the interfaces to choose from",
	Run: func(cmd *cobra.Command, args []string) {
		err := listInterfaces()

		if err != nil {
			_ = fmt.Errorf("something went wrong getting interface list")
		}
	},
}

// ListInterfaces prints available network interfaces.
func listInterfaces() error {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		return err
	}

	for _, device := range devices {
		fmt.Printf("Name: %s, Description: %s\n", device.Name, device.Description)
	}

	return nil
}
