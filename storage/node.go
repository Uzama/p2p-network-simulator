package storage

type node struct {
	peer  *peer
	left  *node
	right *node
}

func newNode(peer *peer) *node {
	return &node{
		peer: peer,
	}
}
