package ip

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

type IPHeader []byte

func (ip IPHeader) GetVersion() uint8 {
	mask := byte(0xf0)
	b := byte(ip[0])
	b = mask & b
	b = (b >> 4)
	return b
}

func (ip IPHeader) GetIHL() uint8 {
	mask := byte(0x0f)
	b := byte(ip[0])
	b = mask & b
	return b
}

func (ip IPHeader) GetDSCP() uint8 {
	mask := byte(0xf8)
	b := byte(ip[1])
	b = mask & b
	b = (b >> 3)
	return b
}

func (ip IPHeader) GetECN() uint8 {
	mask := byte(0x07)
	b := byte(ip[1])
	b = mask & b
	return b
}

func (ip IPHeader) GetTotalLength() uint16 {
	return binary.BigEndian.Uint16(ip[2:4])
}

func (ip IPHeader) GetIdentification() uint16 {
	return binary.BigEndian.Uint16(ip[4:6])
}

func (ip IPHeader) GetFlags() uint8 {
	mask := byte(0xe0)
	b := byte(ip[7])
	b = mask & b
	b = (b >> 5)
	return b
}

func (ip IPHeader) GetFragmentOffset() uint16 {
	mask := uint16(0x1fff)
	b := binary.BigEndian.Uint16(ip[6:8])
	b = mask & b
	return b
}

func (ip IPHeader) GetTTL() uint8 {
	return uint8(ip[8])
}

func (ip IPHeader) GetProtocol() uint8 {
	return uint8(ip[9])
}

func (ip IPHeader) GetChecksum() uint16 {
	return binary.BigEndian.Uint16(ip[10:12])
}

func (ip IPHeader) GetSourceIP() string {
	return ip.addrToString(12, 16)
}

func (ip IPHeader) SetSourceIP(s string) {
	buf := ip.stringToByteArray(s)
	copy(ip[12:16], buf)
}

func (ip IPHeader) GetDestinationIP() string {
	return ip.addrToString(16, 20)
}

func (ip IPHeader) SetDestinationIP(d string) {
	buf := ip.stringToByteArray(d)
	copy(ip[16:20], buf)
}

func (ip IPHeader) GetPayloadOffset() uint16 {
	thl := ip.GetIHL()
	return uint16(thl) * 4
}

/* Unexported utility Methods */
func (ip IPHeader) addrToString(s, e int) string {
	i := binary.BigEndian.Uint32(ip[s:e])
	ibuf := make([]byte, 4)
	binary.BigEndian.PutUint32(ibuf, i)
	addr := fmt.Sprintf("%d.%d.%d.%d", ibuf[0], ibuf[1], ibuf[2], ibuf[3])
	return addr
}

func (ip IPHeader) stringToByteArray(addr string) []byte {
	sbuf := strings.Split(addr, ".")
	ibuf := make([]byte, 4)
	a, _ := strconv.Atoi(sbuf[0])
	b, _ := strconv.Atoi(sbuf[1])
	c, _ := strconv.Atoi(sbuf[2])
	d, _ := strconv.Atoi(sbuf[3])
	ibuf[0] = byte(a)
	ibuf[1] = byte(b)
	ibuf[2] = byte(c)
	ibuf[3] = byte(d)
	return ibuf
}
