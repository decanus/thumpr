// Package node contains a HOPR node.
package node

import (
	"github.com/libp2p/go-libp2p-core/peer"
)

// Node implements the HOPR node.
type Node struct {

}

// Start starts a HOPR node.
func (n *Node) Start() {

}

// Stop stops a HOPR node.
func (n *Node) Stop() {

}

/// SendMessage is used to send a message over HOPR.
func (n *Node) SendMessage(msg []byte, destination peer.ID) error {
	return nil
}
