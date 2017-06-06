package tcp

import (
	"encoding/binary"
)

type TCPHeader []byte

func (th TCPHeader) GetSourcePort() uint16 {
	return binary.BigEndian.Uint16(th[0:2])
}

func (th TCPHeader) SetSourcePort(p uint16) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, p)
	copy(th[0:2], buf)
}

func (th TCPHeader) GetDestinationPort() uint16 {
	return binary.BigEndian.Uint16(th[2:4])
}

func (th TCPHeader) SetDestinationPort(p uint16) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, p)
	copy(th[2:4], buf)
}

func (th TCPHeader) GetSeqNumber() uint32 {
	return binary.BigEndian.Uint32(th[4:8])
}

func (th TCPHeader) SetSeqNumber(sn uint32) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, sn)
	copy(th[4:8], buf)
}

func (th TCPHeader) GetAckNumber() uint32 {
	return binary.BigEndian.Uint32(th[8:12])
}

func (th TCPHeader) SetAckNumber(sn uint32) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, sn)
	copy(th[8:12], buf)
}

func (th TCPHeader) GetDataOffset() uint8 {
	mask := byte(0xf0)
	b := byte(th[12])
	b = mask & b
	b = (b >> 4)
	return b
}

func (th TCPHeader) GetReserved() uint8 {
	mask := byte(0x7)
	b := byte(th[12])
	b = mask & b
	return b
}

func (th TCPHeader) GetNS() bool {
	mask := byte(0x1)
	b := byte(th[12])
	b = mask & b
	return booleanValue(b)
}

func (th TCPHeader) GetCWR() bool {
	mask := byte(0x80)
	b := byte(th[13])
	b = mask & b
	b = (b >> 7)
	return booleanValue(b)
}

func (th TCPHeader) GetECE() bool {
	mask := byte(0x40)
	b := byte(th[13])
	b = mask & b
	b = (b >> 6)
	return booleanValue(b)
}

func (th TCPHeader) GetURG() bool {
	mask := byte(0x20)
	b := byte(th[13])
	b = mask & b
	b = (b >> 5)
	return booleanValue(b)
}

func (th TCPHeader) GetACK() bool {
	mask := byte(0x10)
	b := byte(th[13])
	b = mask & b
	b = (b >> 4)
	return booleanValue(b)
}

func (th TCPHeader) SetACK(v bool) {
	mask := byte(0x10)
	b := byte(th[13])
	if v {
		b = mask | b
	} else {
		mask = ^mask
		b = mask & b
	}
	th[13] = b
}

func (th TCPHeader) GetPSH() bool {
	mask := byte(0x8)
	b := byte(th[13])
	b = mask & b
	b = (b >> 3)
	return booleanValue(b)
}

func (th TCPHeader) GetRST() bool {
	mask := byte(0x4)
	b := byte(th[13])
	b = mask & b
	b = (b >> 2)
	return booleanValue(b)
}

func (th TCPHeader) GetSYN() bool {
	mask := byte(0x2)
	b := byte(th[13])
	b = mask & b
	b = (b >> 1)
	return booleanValue(b)
}

func (th TCPHeader) GetFIN() bool {
	mask := byte(0x1)
	b := byte(th[13])
	b = mask & b
	return booleanValue(b)
}

func (th TCPHeader) SetFIN(v bool) {
	mask := byte(0x1)
	b := byte(th[13])
	if v {
		b = mask | b
	} else {
		mask = ^mask
		b = mask & b
	}
	th[13] = b
}

func (th TCPHeader) GetWindowSize() uint16 {
	return binary.BigEndian.Uint16(th[14:16])
}

func (th TCPHeader) SetWindowSize(s uint16) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, s)
	copy(th[14:16], buf)
}

func (th TCPHeader) GetChecksum() uint16 {
	return binary.BigEndian.Uint16(th[16:18])
}

func (th TCPHeader) SetChecksum(c uint16) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, c)
	copy(th[16:18], buf)
}

func (th TCPHeader) GetUrgentPointer() uint16 {
	return binary.BigEndian.Uint16(th[18:20])
}

func (th TCPHeader) SetUrgentPointer(u uint16) {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, u)
	copy(th[18:20], buf)
}

func (th TCPHeader) GetPayloadOffset() uint16 {
	do := th.GetDataOffset()
	return uint16(do) * 4
}

/* Utility Methods */

func booleanValue(b byte) bool {
	if b > byte(0) {
		return true
	} else {
		return false
	}
}
