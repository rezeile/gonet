package main

import (
	"fmt"
	"github.com/rezeile/gonet/ip"
	"github.com/rezeile/gonet/tcp"
	"github.com/rezeile/gonet/tun"
	"log"
)

func handleConnection(conn *tcp.TCPConn) {
	fmt.Println("TCP Connection Established")
	for {
		ih := <-conn.Reader
		ip.PrintIPHeader(ih)
	}
}

func main() {
	tun.Configure()
	ln, err := tcp.Listen("10.0.0.2", 23)
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
