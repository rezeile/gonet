package main

import (
	"github.com/rezeile/gonet/debug"
	"github.com/rezeile/gonet/ip"
	"github.com/rezeile/gonet/udp"
	"github.com/songgao/water"
	"log"
)

func echoMessage(ifce *water.Interface, packet []byte) {
	var ih ip.IPHeader = packet
	debug.PrintIPHeader(ih)
	var uh udp.UDPHeader = packet[ih.GetPayloadOffset():]
	debug.PrintUDPHeader(uh)
	/* Rewrite Packets */
	sip := ih.GetSourceIP()
	dip := ih.GetDestinationIP()
	ih.SetSourceIP(dip)
	ih.SetDestinationIP(sip)
	sport := uh.GetSourcePort()
	uh.SetSourcePort(uh.GetDestinationPort())
	uh.SetDestinationPort(sport)

	/* Write To Interface */
	ifce.Write(ih)
	/* Log New Results */
	debug.PrintIPHeader(ih)
	debug.PrintUDPHeader(uh)
}

func main() {
	/* Create New TUN interface */
	ifce, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Interface  Name: %s\n", ifce.Name())
	packet := make([]byte, 2000)
	for {
		n, err := ifce.Read(packet)
		if err != nil {
			log.Fatal(err)
		}
		echoMessage(ifce, packet[:n])
	}
}
