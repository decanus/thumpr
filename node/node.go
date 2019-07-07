// Package node contains a HOPR node.
package node

import (
	"github.com/libp2p/go-libp2p-core/peer"
)

// Node implements the HOPR node.
type Node struct {

}

/// SendMessage is used to send a message over HOPR.
func (n *Node) SendMessage(msg []byte, destination peer.ID) error {
	return nil
}
