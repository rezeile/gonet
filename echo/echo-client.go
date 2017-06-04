package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	/* Resolve destination TCPAddr */
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:3030")
	if err != nil {
		fmt.Println(err)
		return
	}
	/* Dial the destination host */
	tcpconn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println(err)
		if tcpconn != nil {
			tcpconn.Close()
		}
		return
	}
	/* Send command line arguments to server and wait for response */
	args := os.Args[1:]
	str := []byte(strings.Join(args, " "))

	_, err = tcpconn.Write(str)
	if err != nil {
		fmt.Println("Error while writing")
		return
	}

	buf := make([]byte, 4096)
	tcpconn.Read(buf)
	fmt.Println(string(buf))
}
