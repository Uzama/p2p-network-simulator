package storage

type tree struct {
	root *peer
}

func newTree(peer *peer) *tree {
	return &tree{
		root: peer,
	}
}

func (t *tree) locate(id int) *peer {
	if t.root == nil {
		return nil
	}

	queue := make([]*peer, 0)

	queue = append(queue, t.root)

	for len(queue) != 0 {

		current := queue[0]
		queue = queue[1:]

		if current.id == id {
			return current
		}

		if len(current.children) > 0 {
			queue = append(queue, current.children...)
		}
	}

	return nil
}
