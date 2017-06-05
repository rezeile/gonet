package tcp

import (
	"encoding/binary"
)

/* TCP Header Fields
srcPort   uint16
dstPort   uint16
seqNum    uint32
ackNum    uint32
do_res_cp uint16 (data offset, reserve, and control bits)
window    uint16
checksum  uint16
urgPtr    uint16
options   uint32
payload   []byte*/

/* Expectes a TCP Header as byte stream []byte */
type TCPHeader []byte

func (th TCPHeader) GetSrcPort() uint16 {
	return binary.BigEndian.Uint16(th[0:2])
}

func (th TCPHeader) SetSrcPort(srcPort uint16) {
	binary.BigEndian.PutUint16(th[0:2], srcPort)
}

func (th TCPHeader) GetDstPort() uint16 {
	return binary.BigEndian.Uint16(th[2:4])
}

func (th TCPHeader) SetDstPort(srcPort uint16) {
	binary.BigEndian.PutUint16(th[2:4], srcPort)
}

/*func (th *TCPHeader) GetSrcPort() uint16 {
	return th.srcPort
}

func (th TCPHeader) {

}
func (th *TCPHeader) GetDstPort() uint16 {
	return th.dstPort
}*/

/*func ParseTCPHeader(packet []byte) *TCPHeader {
	th := &TCPHeader{}
	th.srcPort = binary.BigEndian.Uint16(packet[0:2])
	th.dstPort = binary.BigEndian.Uint16(packet[2:4])
	th.seqNum = binary.BigEndian.Uint32(packet[4:8])
	th.ackNum = binary.BigEndian.Uint32(packet[8:12])
	th.drc = binary.BigEndian.Uint16(packet[12:14])
	th.window = binary.BigEndian.Uint16(packet[14:16])
	th.checksum = binary.BigEndian.Uint16(packet[16:18])
	th.urgPtr = binary.BigEndian.Uint16(packet[18:20])
	th.options = binary.BigEndian.Uint32(packet[20:24])
	th.payload = packet[24:]
	return th
}*/
