package debug

import (
	"fmt"
	"github.com/rezeile/gonet/ip"
	"github.com/rezeile/gonet/tcp"
	"github.com/rezeile/gonet/udp"
)

func PrintIPHeader(ip ip.IPHeader) {
	fmt.Printf("IP Header\n")
	fmt.Printf("-----------------------------\n")
	fmt.Printf("Version: %d\n", ip.GetVersion())
	fmt.Printf("IHL: %d\n", ip.GetIHL())
	fmt.Printf("DSCP: %#x\n", ip.GetDSCP())
	fmt.Printf("ECN: %#x\n", ip.GetECN())
	fmt.Printf("Total Length: %d\n", ip.GetTotalLength())
	fmt.Printf("Identification: %#x\n", ip.GetIdentification())
	fmt.Printf("Flags: %d\n", ip.GetFlags())
	fmt.Printf("Fragement Offset: %d\n", ip.GetFragmentOffset())
	fmt.Printf("TTL: %d\n", ip.GetTTL())
	fmt.Printf("Protocol: %d\n", ip.GetProtocol())
	fmt.Printf("Checksum: %#x\n", ip.GetChecksum())
	fmt.Printf("Source IP: %s\n", ip.GetSourceIP())
	fmt.Printf("Destination IP: %s\n", ip.GetDestinationIP())
	fmt.Printf("-----------------------------\n")
}

func PrintUDPHeader(uh udp.UDPHeader) {
	fmt.Printf("UDP Header\n")
	fmt.Printf("-----------------------------\n")
	fmt.Printf("Source Port: %d\n", uh.GetSourcePort())
	fmt.Printf("Destination Port: %d\n", uh.GetDestinationPort())
	fmt.Printf("Length: %d\n", uh.GetLength())
	fmt.Printf("Checksum: %#x\n", uh.GetChecksum())
	fmt.Printf("-----------------------------\n")
}

func tcpControlBits(th tcp.TCPHeader) string {
	var ns, cwr, ece, urg, ack, psh, rst, syn, fin string = "NO", "NO", "NO", "NO", "NO", "NO", "NO", "NO", "NO"
	if th.GetNS() > 0 {
		ns = "YES"
	}
	if th.GetCWR() > 0 {
		cwr = "YES"
	}
	if th.GetECE() > 0 {
		ece = "YES"
	}
	if th.GetURG() > 0 {
		urg = "YES"
	}
	if th.GetACK() > 0 {
		ack = "YES"
	}
	if th.GetPSH() > 0 {
		psh = "YES"
	}
	if th.GetRST() > 0 {
		rst = "YES"
	}
	if th.GetSYN() > 0 {
		syn = "YES"
	}
	if th.GetFIN() > 0 {
		fin = "YES"
	}
	return fmt.Sprintf("[NS: %s],[CWR: %s],[ECE: %s],[URG: %s],[ACK: %s],[PSH: %s],[RST: %s],[SYN: %s],[FIN: %s]", ns, cwr, ece, urg, ack, psh, rst, syn, fin)

}

func PrintTCPHeader(th tcp.TCPHeader) {
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
