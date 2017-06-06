package ip

import (
	"github.com/songgao/water"
	"log"
)

const TUN_MAX = 8

var IPTun *IP

type IP struct {
	ifce   *water.Interface
	Writer chan IPHeader
}

func NetworkTunnel() (*IP, error) {
	i, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		log.Fatal(err)
	}
	return &IP{
		ifce:   i,
		Writer: make(chan IPHeader, TUN_MAX),
	}, err
}

func (i *IP) Name() string {
	return i.ifce.Name()
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
