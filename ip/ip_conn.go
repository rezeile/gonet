package ip

import (
	"fmt"
	"github.com/songgao/water"
	"log"
)

type IP struct {
	ifce *water.Interface
}

func NetworkTunnel() (*IP, error) {
	ifce, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Interface %s\n", ifce.Name())
	return &IP{ifce}, err
}

func (i *IP) Read(b []byte) (int, error) {
	return i.ifce.Read(b)
}

func (i *IP) Write(b []byte) (int, error) {
	return i.ifce.Write(b)
}

func (i *IP) Close() error {
	return i.ifce.Close()
}
