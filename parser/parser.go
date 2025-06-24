package parser

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// ParseAndPrintPacket extracts and prints protocol details
func ParseAndPrintPacket(packet gopacket.Packet) {
	// Ethernet layer
	if ethLayer := packet.Layer(layers.LayerTypeEthernet); ethLayer != nil {
		eth, _ := ethLayer.(*layers.Ethernet)
		fmt.Printf("Ethernet: %s -> %s | Type: %s\n", eth.SrcMAC, eth.DstMAC, eth.EthernetType)
	}

	// IPv4 Layer
	if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
		ip, _ := ipLayer.(*layers.IPv4)
		fmt.Printf("IPv4: %s -> %s | Protocol: %s\n", ip.SrcIP, ip.DstIP, ip.Protocol)
	}

	// TCP Layer
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		tcp, _ := tcpLayer.(*layers.TCP)
		fmt.Printf("TCP: %d -> %d | SYN: %t, ACK: %t\n", tcp.SrcPort, tcp.DstPort, tcp.SYN, tcp.ACK)
		fmt.Printf("Data:\n%s", packet.Data())
	}

	// UDP Layer
	if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
		udp, _ := udpLayer.(*layers.UDP)
		fmt.Printf("UDP: %d -> %d\n", udp.SrcPort, udp.DstPort)
	}

	if dnsLayer := packet.Layer(layers.LayerTypeDNS); dnsLayer != nil {
		dns, _ := dnsLayer.(*layers.DNS)
		fmt.Printf("DNS: ID=%d, Qs=%d, Ans=%d\n", dns.ID, len(dns.Questions), len(dns.Answers))
		for _, q := range dns.Questions {
			fmt.Printf("  Query: %s (%s)\n", string(q.Name), q.Type)
		}
	}

	fmt.Println("---")
}
