package main

import (
	"github.com/rezeile/gonet/tcp"
	"log"
)

func testTCP() {
	l, err := tcp.Listen("10.0.0.2", 23)
	if err != nil {
		log.Fatal(err)
	}
	l.Accept()
}

func main() {
	testTCP()
}
