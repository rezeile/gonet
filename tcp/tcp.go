package mytcp

import (
	"fmt"
)

type TCPHeader struct {
	srcPort  uint16
	dstPort  uint16
	seqNum   uint32
	ackNum   uint32
	drc      uint16 /* Data Offsets, Reserved, and Control Bits */
	window   uint16
	checksum uint16
	urgPtr   uint16
	options  uint32
	payload  []byte
}
