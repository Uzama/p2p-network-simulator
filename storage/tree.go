package storage

type tree struct {
	root *peer
}

func newTree(peer *peer) *tree {
	return &tree{
		root: peer,
	}
}
