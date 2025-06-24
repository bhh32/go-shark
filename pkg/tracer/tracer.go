package tracer

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func RunTrace(dest string, maxHops int) error {
	ipAddr, err := net.ResolveIPAddr("ip4", dest)
	if err != nil {
		return fmt.Errorf("could not resolve %s: %v", dest, err)
	}

	fmt.Printf("Trace route to %s (%s), %d hops max\n", dest, ipAddr.String(), maxHops)

	// Listen for icmp replies
	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")

	if err != nil {
		return fmt.Errorf("listen error: %v", err)
	}

	defer c.Close()

	basePort := 33434

	for ttl := 1; ttl <= maxHops; ttl++ {
		dstPort := basePort + rand.Intn(100)
		udpAddr := &net.UDPAddr{IP: ipAddr.IP, Port: dstPort}
		udpCon, err := net.DialUDP("udp", nil, udpAddr)
		if err != nil {
			fmt.Printf("%2d *\n", ttl)
			continue
		}

		pconn := ipv4.NewConn(udpCon)
		if err := pconn.SetTTL(ttl); err != nil {
			udpCon.Close()
			fmt.Printf("%2d *\n", ttl)
			continue
		}

		start := time.Now()
		_, err = udpCon.Write([]byte("GO-SHARK-TRACER"))
		udpCon.Close()

		if err != nil {
			fmt.Printf("%2d \n", ttl)
			continue
		}

		c.SetReadDeadline(time.Now().Add(3 * time.Second))
		reply := make([]byte, 1500)
		n, peer, err := c.ReadFrom(reply)
		elapsed := time.Since(start)

		if err != nil {
			fmt.Printf("%2d %s %.3f ms\n", ttl, peer, float64(elapsed.Microseconds())/1000)
			return nil
		}

		rm, err := icmp.ParseMessage(1, reply[:n])
		if err != nil {
			fmt.Printf("%2d *\n", ttl)
			continue
		}

		switch rm.Type {
		case ipv4.ICMPTypeTimeExceeded:
			fmt.Printf("%2d %s %.3f ms\n", ttl, peer, float64(elapsed.Microseconds())/1000)
		case ipv4.ICMPTypeDestinationUnreachable:
			fmt.Printf("%2d %s %.3f ms (destination reached\n", ttl, peer, float64(elapsed.Microseconds())/1000)
			return nil
		default:
			fmt.Printf("%2d %s [unknown ICMP type]\n", ttl, peer)
		}

		time.Sleep(100 * time.Millisecond)
	}

	return nil
}
