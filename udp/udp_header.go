package udp

import (
	"encoding/binary"
)

type UDPHeader []byte

const UDP_HDR_SIZE = 8

func (uh UDPHeader) GetSourcePort() uint16 {
	res := binary.BigEndian.Uint16(uh[0:2])
	return res
}

func (uh UDPHeader) SetSourcePort(p uint16) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, p)
	copy(uh[0:2], buf)
}

func (uh UDPHeader) GetDestinationPort() uint16 {
	return binary.BigEndian.Uint16(uh[2:4])
}

func (uh UDPHeader) SetDestinationPort(p uint16) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, p)
	copy(uh[2:4], buf)
}

func (uh UDPHeader) GetLength() uint16 {
	return binary.BigEndian.Uint16(uh[4:6])
}

func (uh UDPHeader) SetLength(l uint16) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, l)
	copy(uh[4:6], buf)
}

func (uh UDPHeader) GetChecksum() uint16 {
	return binary.BigEndian.Uint16(uh[6:8])
}

func (uh UDPHeader) SetChecksum(c uint16) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, c)
	copy(uh[6:8], buf)

}

func (uh UDPHeader) GetPayloadOffset() uint16 {
	return UDP_HDR_SIZE
}
