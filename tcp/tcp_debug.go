package tcp

import (
	"fmt"
)

func tcpControlBits(th TCPHeader) string {
	var ns, cwr, ece, urg, ack, psh, rst, syn, fin string = "NO", "NO", "NO", "NO", "NO", "NO", "NO", "NO", "NO"
	if th.GetNS() {
		ns = "YES"
	}
	if th.GetCWR() {
		cwr = "YES"
	}
	if th.GetECE() {
		ece = "YES"
	}
	if th.GetURG() {
		urg = "YES"
	}
	if th.GetACK() {
		ack = "YES"
	}
	if th.GetPSH() {
		psh = "YES"
	}
	if th.GetRST() {
		rst = "YES"
	}
	if th.GetSYN() {
		syn = "YES"
	}
	if th.GetFIN() {
		fin = "YES"
	}
	return fmt.Sprintf("[NS: %s],[CWR: %s],[ECE: %s],[URG: %s],[ACK: %s],[PSH: %s],[RST: %s],[SYN: %s],[FIN: %s]", ns, cwr, ece, urg, ack, psh, rst, syn, fin)

}

func PrintTCPHeader(th TCPHeader) {
	fmt.Printf("TCP Header\n")
	fmt.Printf("-----------------------------\n")
	fmt.Printf("Source Port: %d\n", th.GetSourcePort())
	fmt.Printf("Destination Port: %d\n", th.GetDestinationPort())
	fmt.Printf("Seq Number:  %d\n", th.GetSeqNumber())
	fmt.Printf("Ack Number: %d\n", th.GetAckNumber())
	fmt.Printf("Data Offset: %d\n", th.GetDataOffset())
	fmt.Printf("Reserved: %#x\n", th.GetReserved())
	fmt.Printf("ControlBits: %s\n", tcpControlBits(th))
	fmt.Printf("Window Size: %d\n", th.GetWindowSize())
	fmt.Printf("Checksum: %#x\n", th.GetChecksum())
	fmt.Printf("Urgent Pointer: %d\n", th.GetUrgentPointer())
	fmt.Printf("-----------------------------\n")
}
