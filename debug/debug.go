package debug

import (
	"fmt"
	"github.com/rezeile/gonet/ip"
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
