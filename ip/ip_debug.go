package ip

import (
	"fmt"
)

func PrintIPHeader(ip IPHeader) {
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
