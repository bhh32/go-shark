package utils

import (
	"fmt"

	"github.com/google/gopacket/pcap"
)

// ListInterfaces prints available network interfaces.
func ListInterfaces() error {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		return err
	}

	for _, device := range devices {
		fmt.Printf("Name: %s, Description: %s\n", device.Name, device.Description)
	}

	return nil
}
