// Package node contains a HOPR node.
package node

import (
	"context"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
)

// Node implements the HOPR node.
type Node struct {
	host host.Host
}

func NewNode() *Node {
	return &Node{}
}

// Start starts a HOPR node.
func (n *Node) Start() error {
	h, err := libp2p.New(context.Background())

	if err != nil {
		return err
	}

	n.host = h

	// @todo start reading and writing
	// @todo maybe add status? like on off

	return nil
}

// Stop stops a HOPR node.
func (n *Node) Stop() error {
	if n.host == nil {
		return nil
	}

	err := n.host.Close()
	if err != nil {
		return err
	}
}

// Crawl crawls ... for ... .
func (n *Node) Crawl() {

}

/// SendMessage is used to send a message over HOPR.
func (n *Node) SendMessage(msg []byte, destination peer.ID) error {
	return nil
}
