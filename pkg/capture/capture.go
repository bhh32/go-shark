package capture

import (
	"fmt"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"

	"github.com/bhh32/go-shark/parser"
)

// CapturePackets captures packets on the specified interface
func CapturePackets(iface string, durationSecs int, filter string) error {
	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)

	if err != nil {
		return fmt.Errorf("could not open device %s: %v", iface, err)
	}

	defer handle.Close()

	if filter != "" {
		if err := handle.SetBPFFilter(filter); err != nil {
			return fmt.Errorf("could not set BPF filter: %v", err)
		}

		fmt.Println("Applied filter: ", filter)
	}

	packetSrc := gopacket.NewPacketSource(handle, handle.LinkType())
	timeout := time.After(time.Duration(durationSecs) * time.Second)

	fmt.Printf("Capturing packets on %s for %d seconds...\n", iface, durationSecs)

	for {
		select {
		case packet := <-packetSrc.Packets():
			if packet == nil {
				continue
			}

			// Parse and print the packet information
			parser.ParseAndPrintPacket(packet)
		case <-timeout:
			fmt.Println("Capture complete.")
			return nil
		}
	}
}
