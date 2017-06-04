package main

import (
	"fmt"
	"github.com/rezeile/goudp/udp"
	"github.com/songgao/water"
	"github.com/songgao/water/waterutil"
	"log"
)

func echoMessage(ifce *water.Interface, packet *[]byte) {
	uh := udp.ParseUDPHeader(waterutil.IPv4Payload(*packet))
	/* Get Source IP */
	sip := waterutil.IPv4Source(*packet)
	sport := uh.GetSrcPort()
	fmt.Println("Source IP: ", sip.String())
	fmt.Println("Source Port: ", sport)
	/* Get Destination IP */
	dip := waterutil.IPv4Destination(*packet)
	dport := uh.GetDstPort()
	fmt.Println("Destination IP: ", dip.String())
	fmt.Println("Destination Port: ", dport)
	fmt.Println("----------------------------------")

	/* Rewrite packet */
	waterutil.SetIPv4Source(*packet, dip)
	waterutil.SetIPv4Destination(*packet, sip)
	uh.SetSrcPort(*packet, dport)
	uh.SetDstPort(*packet, sport)

	/* Print actual message */
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
		ifce.Write(packet[:n])
	}
}
