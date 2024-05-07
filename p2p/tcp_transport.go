package p2p

import (
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listner       net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *Transport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}

}

func Tset() {
	t := NewTCPTransport(":4344").(*TCPTransport)

	t.listner.Accept()
}
