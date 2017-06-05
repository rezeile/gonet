package tcp

type TCPListener struct {
}

func Listen(ipaddr string, port uint16) (*TCPListener, error) {
	return &TCPListener{}, nil
}

func (l *TCPListener) Accept() (*TCPConn, error) {
	return &TCPConn{}, nil
}

func (l *TCPListener) Close() error {
	return nil
}

func (l *TCPListener) Address() string {
	return ""
}
