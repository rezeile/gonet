package udp

import (
	"encoding/binary"
)

type UDPHeader struct {
	srcPort  uint16
	dstPort  uint16
	length   uint16
	checksum uint16
	payload  []byte
}

func (uh *UDPHeader) GetSrcPort() uint16 {
	return uh.srcPort
}

func (uh *UDPHeader) SetSrcPort(packet []byte, srcPort uint16) {
	binary.BigEndian.PutUint16(packet[0:2], srcPort)
	uh.srcPort = srcPort
}

func (uh *UDPHeader) GetDstPort() uint16 {
	return uh.dstPort
}

func (uh *UDPHeader) SetDstPort(packet []byte, dstPort uint16) {
	binary.BigEndian.PutUint16(packet[2:4], dstPort)
	uh.dstPort = dstPort
}

func (uh *UDPHeader) GetPayload() []byte {
	return uh.payload
}

func ParseUDPHeader(packet []byte) *UDPHeader {
	uh := &UDPHeader{}
	uh.srcPort = binary.BigEndian.Uint16(packet[0:2])
	uh.dstPort = binary.BigEndian.Uint16(packet[2:4])
	uh.checksum = binary.BigEndian.Uint16(packet[16:18])
	uh.payload = packet[24:]
	return uh
}
