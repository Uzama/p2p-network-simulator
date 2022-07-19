package treap

import "p2p-network-simulator/storage/tree"

// treap node
type node struct {
	peer  *tree.Peer // use Id as key and CurrentCapacity as priority
	left  *node
	right *node
}

// newNode: create new node
func newNode(peer *tree.Peer) *node {
	return &node{
		peer: peer,
	}
}

// get: return the peer
func (n *node) get() *tree.Peer {
	return n.peer
}
