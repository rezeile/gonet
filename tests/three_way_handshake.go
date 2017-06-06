package main

import (
	"fmt"
	"github.com/rezeile/gonet/tcp"
	"log"
)

func testTCP() {
	//  TODO: Dynamically Configure TUN Interface */
	l, err := tcp.Listen("10.0.0.2", 23)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
	}
}

func main() {
	testTCP()
}
