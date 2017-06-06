package udp

func PrintUDPHeader(uh UDPHeader) {
	fmt.Printf("UDP Header\n")
	fmt.Printf("-----------------------------\n")
	fmt.Printf("Source Port: %d\n", uh.GetSourcePort())
	fmt.Printf("Destination Port: %d\n", uh.GetDestinationPort())
	fmt.Printf("Length: %d\n", uh.GetLength())
	fmt.Printf("Checksum: %#x\n", uh.GetChecksum())
	fmt.Printf("-----------------------------\n")
}
