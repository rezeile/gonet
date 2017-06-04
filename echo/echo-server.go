package main

import (
	"fmt"
	"net"
)

const CHUNK_SIZE = 4096

func handleTCPConnection(conn *net.TCPConn) {
	/* Read from the connection */
	buff := make([]byte, CHUNK_SIZE)
	for {
		n, err := conn.Read(buff)
		/* If everything has been read or error while reading, continue */
		if n == 0 || err != nil {
			continue
		}
		/* Print the received message locally, and send it back */
		fmt.Print(string(buff[0:n]))
		conn.Write(buff)
	}
}

func main() {
	/* Resolve tcp address of a socket */
	addr, err := net.ResolveTCPAddr("tcp", ":3030")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Connected to: ", addr.String(), " on ", addr.Network())
	}
	/* Listen to connection requests at tcp address */
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Accepting connections on: ", listener.Addr())
	defer listener.Close()

	/* Continually accept connections */
	for {
		tcpconn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			if tcpconn != nil {
				tcpconn.Close()
			}
			continue
		}
		go handleTCPConnection(tcpconn)
	}
}
