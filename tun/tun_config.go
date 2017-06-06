package tun

import (
	"fmt"
	"github.com/rezeile/gonet/ip"
	"github.com/rezeile/gonet/tcp"
	"log"
)

const MTU = 1500
const MAX_TUN = 8

func Configure() {
	tcp.Listeners = make(map[string]*tcp.TCPListener)
	tcp.Connections = make(map[string]*tcp.TCPConn)
	configureIP()
	go demultiplex()
}

func configureIP() {
	var err error
	ip.IPTun, err = ip.NetworkTunnel()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully Configured Interface %s\n", ip.IPTun.Name())
}

func demultiplex() {
	/* Listen for arriving ip datagrams  at network interface */
	ipt := ip.IPTun
	datagram := make([]byte, MTU)
	for {
		select {
		case ih := <-ipt.Writer:
			ipt.Write([]byte(ih))
		default:
			/* Read arriving packet at interface */
			n, err := ipt.Read(datagram)
			if err != nil {
				log.Fatal(err)
			}
			ih := ip.IPHeader(datagram[:n])
			th := tcp.TCPHeader(ih[ih.GetPayloadOffset():])
			fmt.Println(th)
			if k := getListenerKey(ih, th); k != "" {
				fmt.Println("For Listener")
				ln := tcp.Listeners[k]
				ln.Reader <- ih
			}

			if k := getConnectionKey(ih, th); k != "" {
				fmt.Println("For Connection")
				conn := tcp.Connections[k]
				if conn.GetState() == tcp.ESTABLISHED && th.GetFIN() {
					conn.InitiatePassiveClose(ih)
				}
				if conn.GetState() == tcp.LAST_ACK && th.GetACK() {
					conn.CompletePassiveClose()
				}
			}
		}
	}
}

func getListenerKey(ih ip.IPHeader, th tcp.TCPHeader) string {
	lnkey := tcp.GenerateLKey(ih.GetDestinationIP(), th.GetDestinationPort())
	if ln, ok := tcp.Listeners[lnkey]; ok {
		if ln.GetState() == tcp.LISTEN && th.GetSYN() {
			return lnkey
		}
		if ln.GetState() == tcp.SYN_RECEIVED && th.GetACK() {
			return lnkey
		}
	}
	return ""
}

func getConnectionKey(ih ip.IPHeader, th tcp.TCPHeader) string {
	sip := ih.GetDestinationIP()
	sp := th.GetDestinationPort()
	dip := ih.GetSourceIP()
	dp := th.GetSourcePort()
	ckey := tcp.GenerateCKey(sip, sp, dip, dp)
	if _, ok := tcp.Connections[ckey]; ok {
		return ckey
	}
	return ""
}
