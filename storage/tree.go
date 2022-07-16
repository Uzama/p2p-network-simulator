package storage

type tree struct {
	root *peer
}

func newTree() *tree {
	return &tree{}
}
