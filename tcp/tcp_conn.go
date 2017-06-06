package tcp

import (
	"github.com/rezeile/gonet/ip"
)

var Connections map[string]*TCPConn

type TCPConn struct {
	Writer          chan ip.IPHeader
	Reader          chan ip.IPHeader
	sourceIP        string
	sourcePort      uint16
	destinationIP   string
	destinationPort uint16
	state           uint8
	mss             uint16 /* Maximum Segment Size */
	sendWindow      uint16
	receiveWindow   uint16
	nextAckNumber   uint32
	nextSeqNumber   uint32
}

/* Active (client) TCP Open */
func Dial(ipaddr string, port uint16) (*TCPConn, error) {
	return &TCPConn{}, nil
}

func (c *TCPConn) Read(b []byte) (int, error) {
	return 0, nil
}

func (c *TCPConn) Write(b []byte) (int, error) {
	return 0, nil
}

func (c *TCPConn) RemoteAddr() string {
	return ""
}

func (c *TCPConn) LocalAddr() string {
	return ""
}

func (c *TCPConn) Close() error {
	return nil
}

/* Called when FIN packet is received */
func (c *TCPConn) InitiatePassiveClose(ih ip.IPHeader) {
	/* Send Ack */
	c.Writer <- createAck(ih)

	/* Move to CLOSE_WAIT state */
	c.state = CLOSE_WAIT

	/* Move to LAST_ACK state */
	c.state = LAST_ACK
}

func (c *TCPConn) CompletePassiveClose() {
	/* Free c from Connections list */
	key := GenerateCKey(c.sourceIP, c.sourcePort, c.destinationIP, c.destinationPort)
	close(c.Reader)
	delete(Connections, key)
}

func (c *TCPConn) GetState() uint8 {
	return c.state
}

/* Utility Methods */
func createAck(ih ip.IPHeader) ip.IPHeader {
	/* Extract Return Fields */
	sourceIP := ih.GetSourceIP()
	destinationIP := ih.GetDestinationIP()
	th := TCPHeader(ih[ih.GetPayloadOffset():])
	sourcePort := th.GetSourcePort()
	destinationPort := th.GetDestinationPort()

	/* Make copy of ip */
	ipPacket := ip.IPHeader(make([]byte, ih.GetTotalLength()))
	copy([]byte(ipPacket), []byte(ih))
	tcpPacket := TCPHeader(ipPacket[ipPacket.GetPayloadOffset():])

	/* Populate Fields for ACK packet */
	ipPacket.SetSourceIP(destinationIP)
	ipPacket.SetDestinationIP(sourceIP)
	tcpPacket.SetSourcePort(destinationPort)
	tcpPacket.SetDestinationPort(sourcePort)
	tcpPacket.SetACK(true)
	tcpPacket.SetAckNumber(th.GetSeqNumber() + 1)
	tcpPacket.SetChecksum(ComputeTCPChecksum(ipPacket))
	return ipPacket
}

func createFin(ih ip.IPHeader) ip.IPHeader {
	/* Extract Return Fields */
	sourceIP := ih.GetSourceIP()
	destinationIP := ih.GetDestinationIP()
	th := TCPHeader(ih[ih.GetPayloadOffset():])
	sourcePort := th.GetSourcePort()
	destinationPort := th.GetDestinationPort()

	/* Make copy of ip */
	ipPacket := ip.IPHeader(make([]byte, ih.GetTotalLength()))
	copy([]byte(ipPacket), []byte(ih))
	tcpPacket := TCPHeader(ipPacket[ipPacket.GetPayloadOffset():])

	/* Populate Fields for FIN packet */
	ipPacket.SetSourceIP(destinationIP)
	ipPacket.SetDestinationIP(sourceIP)
	tcpPacket.SetSourcePort(destinationPort)
	tcpPacket.SetDestinationPort(sourcePort)
	tcpPacket.SetFIN(true)
	tcpPacket.SetACK(false)
	tcpPacket.SetAckNumber(th.GetSeqNumber() + 1)
	tcpPacket.SetChecksum(ComputeTCPChecksum(ipPacket))
	PrintTCPHeader(tcpPacket)
	return ipPacket
}
