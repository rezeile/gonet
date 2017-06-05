package tcp

type TCPConn struct {
	sourceIP        string
	sourcePort      uint16
	destinationIP   string
	destinationPort uint16
	state           uint8
	MSS             uint16 /* Maximum Segment Size */
	sendWindow      uint16
	receiveWindow   uint16
}
