package main

import (
	"fmt"
	"github.com/songgao/water"
	"log"
)

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
		fmt.Printf("%v\n", packet)
	}
}
