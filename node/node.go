// Package node contains a HOPR node.
package node

import (
	"context"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

// Node implements the HOPR node.
type Node struct {
	addr multiaddr.Multiaddr
	pk   crypto.PrivKey

	host host.Host
}

func NewNode(addr multiaddr.Multiaddr, pk crypto.PrivKey) *Node {
	return &Node{
		addr: addr,
		pk:   pk,
	}
}

// Start starts a HOPR node.
func (n *Node) Start() error {
	h, err := libp2p.New(
		context.Background(),
		libp2p.ListenAddrs(n.addr),
		libp2p.Identity(n.pk),
	)

	if err != nil {
		return err
	}

	n.host = h

	// @todo start reading and writing
	// @todo maybe add status? like on off

	return nil
}

// Stop stops a HOPR node.
func (n *Node) Stop() {
	if n.host == nil {
		return
	}

	err := n.host.Close()
	if err != nil {
		panic(err) // @todo
	}
}

// Crawl crawls ... for ... .
func (n *Node) Crawl() {

}

/// SendMessage is used to send a message over HOPR.
func (n *Node) SendMessage(msg []byte, destination peer.ID) error {
	return nil
}
