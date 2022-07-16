package storage

type treap struct {
	root     *node
	tempRoot *node
}

func newTreap() *treap {
	return &treap{}
}

func (t *treap) insert(peer *peer) {
	recursiveInsert(t.root, peer)
}

func (t *treap) delete(id int) {
	recursiveDelete(t.root, id)
}

func (t *treap) mostCapacityPeer() *peer {
	if t.root == nil {
		return nil
	}

	return t.root.getPeer()
}

func rightRotate(root *node) {
	left := root.left
	subTree := left.right

	left.right = root
	root.left = subTree

	root = left
}

func leftRotate(root *node) {
	right := root.right
	subTree := right.left

	right.left = root
	root.right = subTree

	root = right
}

func recursiveInsert(root *node, peer *peer) {
	if root == nil {
		root = newNode(peer)
		return
	}

	if peer.id <= root.peer.id {

		recursiveInsert(root.left, peer)

		heapPropertyViolated := root.left.peer.currentCapacity > root.peer.currentCapacity
		if root.left != nil && heapPropertyViolated {
			rightRotate(root)
		}

		return
	}

	recursiveInsert(root.right, peer)

	heapPropertyViolated := root.right.peer.currentCapacity > root.peer.currentCapacity
	if root.right != nil && heapPropertyViolated {
		leftRotate(root)
	}
}

func recursiveDelete(root *node, id int) {
	if root == nil {
		return
	}

	if id < root.peer.id {
		recursiveDelete(root.left, id)
		return
	}

	if id > root.peer.id {
		recursiveDelete(root.right, id)
		return
	}

	// no children
	if root.left == nil && root.right == nil {
		root = nil
		return
	}

	// having both children
	if root.left != nil && root.right != nil {

		if root.left.peer.currentCapacity < root.right.peer.currentCapacity {
			leftRotate(root)
			recursiveDelete(root.left, id)
			return
		}

		rightRotate(root)
		recursiveDelete(root.right, id)
		return

	}

	// having single child
	if root.right != nil {
		root = root.right
		return
	}

	root = root.left
}
