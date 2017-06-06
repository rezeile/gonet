package tcp

import (
	"bytes"
	"github.com/rezeile/gonet/ip"
	"strconv"
)

const (
	MSS = 1500 /* Maximum Transmission Unit  */
)

var Listeners map[string]*TCPListener

type TCPListener struct {
	Writer          chan ip.IPHeader
	Reader          chan ip.IPHeader
	sourceIP        string
	sourcePort      uint16
	destinationIP   string
	destinationPort uint16
	state           uint8
	mss             uint16
}

func (l *TCPListener) GetState() uint8 {
	return l.state
}

/* Passive TCP Open */
func Listen(ipaddr string, port uint16) (*TCPListener, error) {
	l := &TCPListener{
		Writer:          ip.IPTun.Writer,
		Reader:          make(chan ip.IPHeader),
		sourceIP:        ipaddr,
		sourcePort:      port,
		destinationIP:   "",
		destinationPort: 0,
		state:           LISTEN,
		mss:             MSS}
	/* Add to Listener set */
	Listeners[GenerateLKey(ipaddr, port)] = l
	return l, nil
}

func (l *TCPListener) Accept() (*TCPConn, error) {
	/* Wait for a SYN Packet */
	ih := <-l.Reader
	th := TCPHeader(ih[ih.GetPayloadOffset():])
	if th.GetSYN() {
		l.state = SYN_RECEIVED
		l.Writer <- createSynAck(ih)
	}

	/* Wait for an ACK packet after sending SYN Ack  */
	ih = <-l.Reader
	th = TCPHeader(ih[ih.GetPayloadOffset():])
	if th.GetACK() {
		/* Restore l.state */
		l.state = LISTEN
		/* Return new connection */
		c := &TCPConn{
			Writer:          ip.IPTun.Writer,
			Reader:          make(chan ip.IPHeader),
			sourceIP:        ih.GetDestinationIP(),
			sourcePort:      th.GetDestinationPort(),
			destinationIP:   ih.GetSourceIP(),
			destinationPort: th.GetSourcePort(),
			state:           ESTABLISHED,
			mss:             1500,
			sendWindow:      0,
			receiveWindow:   0,
			nextAckNumber:   th.GetSeqNumber() + 1,
			nextSeqNumber:   th.GetAckNumber() + 1,
		}
		/* Add to Connections */
		addToConnectionList(c)
		/* Return New Connection */
		return c, nil
	}
	/* Return TCP Connection  */
	return &TCPConn{}, nil
}

func (l *TCPListener) Close() error {
	return nil
}

func (l *TCPListener) Address() string {
	return ""
}

/* Three Way Handshake Utility Methods */
func GenerateLKey(ipaddr string, port uint16) string {
	var key bytes.Buffer
	key.WriteString(ipaddr)
	return key.String()
}

func GenerateCKey(s string, sp uint16, d string, dp uint16) string {
	var key bytes.Buffer
	key.WriteString(s)
	key.WriteString(strconv.Itoa(int(sp)))
	key.WriteString(d)
	return key.String()
}

func addToConnectionList(c *TCPConn) {
	key := GenerateCKey(c.sourceIP, c.sourcePort, c.destinationIP, c.destinationPort)
	Connections[key] = c
}

func createSynAck(ih ip.IPHeader) ip.IPHeader {
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

	/* Populate Fields for SYN ACK packet */
	ipPacket.SetSourceIP(destinationIP)
	ipPacket.SetDestinationIP(sourceIP)
	tcpPacket.SetSourcePort(destinationPort)
	tcpPacket.SetDestinationPort(sourcePort)
	tcpPacket.SetACK(true)
	tcpPacket.SetAckNumber(th.GetSeqNumber() + 1)
	tcpPacket.SetChecksum(ComputeTCPChecksum(ipPacket))
	return ipPacket
}
