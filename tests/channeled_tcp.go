package main

import (
	"fmt"
	"github.com/rezeile/gonet/tcp"
	"github.com/rezeile/gonet/tun"
	"log"
)

func handleConnection(conn *tcp.TCPConn) {
	fmt.Println("TCP Connection Established")
}

func main() {
	tun.Configure()
	ln, err := tcp.Listen("10.0.0.5", 23)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
		}

		go handleConnection(conn)
	}
}
