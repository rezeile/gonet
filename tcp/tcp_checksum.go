package tcp

import (
	"encoding/binary"
	"github.com/rezeile/gonet/ip"
)

func generateTCPPseudoHeader(ih ip.IPHeader) []byte {
	pseudoHeader := make([]byte, 12)
	/* Source Address */
	copy(pseudoHeader[0:4], ih[12:16])
	/* Destination Address */
	copy(pseudoHeader[4:8], ih[16:20])
	/* Zeros */
	pseudoHeader[8] = byte(0)
	/* Protocol */
	pseudoHeader[9] = ih[9]
	/* TCP Length */
	binary.BigEndian.PutUint16(pseudoHeader[10:12], ih.GetTCPLength())
	return pseudoHeader
}

func sumPair(a, b uint16) uint16 {
	carrymask := uint32(0xffff0000)
	resmask := uint32(0x0000ffff)
	/* Sum a and b */
	rawres := uint32(a) + uint32(b)
	carry := uint16(((rawres & carrymask) >> 16))
	res := uint16((rawres & resmask))
	return res + carry
}

func onesComplementSum(b []byte, length uint16) uint16 {
	m := binary.BigEndian.Uint16(b[0:2])
	n := binary.BigEndian.Uint16(b[2:4])
	res := sumPair(m, n)
	max := int(length) - 2
	for i := 4; i < max; i += 2 {
		c := binary.BigEndian.Uint16(b[i : i+2])
		res = sumPair(res, c)
	}
	return ^res
}

func ComputeTCPChecksum(ih ip.IPHeader) uint16 {
	/* Create 12 byte pseudo header */
	ph := generateTCPPseudoHeader(ih)
	/* Get stored checksum */
	th := TCPHeader(ih[ih.GetPayloadOffset():])
	thcpy := make([]byte, ih.GetTCPLength())
	copy(thcpy, []byte(th))
	TCPHeader(thcpy).SetChecksum(0)
	/* Prepend TCP Psuedo Header */
	h := append(ph, thcpy...)
	phlen := uint16(len(ph))
	/* Compute the One's Complement Checksum */
	res := onesComplementSum(h, phlen+ih.GetTCPLength())
	return res
}
