package main

import (
	"fmt"
	"github.com/rezeile/gotcp/tcp"
	"github.com/songgao/water"
	"github.com/songgao/water/waterutil"
	"log"
)

func getIPFrame(packet []byte) {
	payload := waterutil.IPv4Payload(packet)
	th := tcp.ParseTCPHeader(payload)
	fmt.Println(th.GetSrcPort(), th.GetDstPort())
}

func main() {
	/* Create New TUN interface */
	ifce, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Interface Name: %s\n", ifce.Name())

	packet := make([]byte, 2000)
	for {
		/*n, err := */ ifce.Read(packet)
		/*if err != nil {
			log.Fatal(err)
		}*/
		getIPFrame(packet)
		//log.Printf("Packet Received: %x\n", packet[:n])
	}
}
