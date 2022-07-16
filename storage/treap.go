package storage

type node struct {
	peer peer
	left *node
	right *node
}

type treap struct {
	Root *node
}

func NewTreap() *treap {
	return &treap{}
}

func (t *treap) insert(peer peer) {

}

func (t *treap) delete(peer peer) {
	
}
