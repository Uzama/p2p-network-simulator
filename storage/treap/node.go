package treap

import "p2p-network-simulator/storage/tree"

type node struct {
	peer  *tree.Peer
	left  *node
	right *node
}

func newNode(peer *tree.Peer) *node {
	return &node{
		peer: peer,
	}
}

func (n *node) getPeer() *tree.Peer {
	return n.peer
}
