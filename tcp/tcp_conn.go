package tcp

type TCPConn struct {
	sourceIP        string
	sourcePort      uint16
	destinationIP   string
	destinationPort uint16
	state           uint8
	MSS             uint16 /* Maximum Segment Size */
	sendWindow      uint16
	receiveWindow   uint16
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
