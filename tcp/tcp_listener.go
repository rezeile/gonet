package tcp

import (
	//"fmt"
	"github.com/rezeile/gonet/ip"
	"log"
)

const (
	MSS = 1500 /* Maximum Transmission Unit  */
)

type TCPListener struct {
	writer          *ip.IP
	reader          *ip.IP
	sourceIP        string
	sourcePort      uint16
	destinationIP   string
	destinationPort uint16
	state           uint8
	mss             uint16
}

/* Passive TCP open */
func Listen(ipaddr string, port uint16) (*TCPListener, error) {
	/* Activate the TUN interface */
	ip, err := ip.NetworkTunnel()
	if err != nil {
		log.Fatal(err)
	}
	return &TCPListener{
		writer:          ip,
		reader:          ip,
		sourceIP:        "",
		sourcePort:      0,
		destinationIP:   "",
		destinationPort: 0,
		state:           LISTEN,
		mss:             MSS}, nil
}

func (l *TCPListener) Accept() (*TCPConn, error) {
	buf := make([]byte, l.mss)
	/* Wait for read from lower layer */
	n, err := l.reader.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	/* Obtain IP layer */
	ih := ip.IPHeader(buf[0:n])
	/* Obtain TCP header */
	th := TCPHeader(buf[ih.GetPayloadOffset():])
	/* Expects a SYN Packet */
	if th.GetSYN() {
		p := createSynAck(ih)
		l.writer.Write([]byte(p))
	}
	/* Expect an ACK Packet */
	n, err = l.reader.Read(buf)
	if err != nil {
		log.Fatal(err)
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

/* Three Way Connection Utility Methods */
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

	/* TODO: Update checksum */
	return ipPacket
}
