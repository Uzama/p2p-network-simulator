package storage

import "fmt"

type treap struct {
	root *node
}

func newTreap() *treap {
	return &treap{
		root: nil,
	}
}

func (t *treap) insert(peer *peer) {
	t.root = recursiveInsert(t.root, peer)
}

func (t *treap) delete(id int) {
	t.root = recursiveDelete(t.root, id)
}

func (t *treap) mostCapacityPeer() *peer {
	if t.root == nil {
		return nil
	}

	return t.root.getPeer()
}

func (t *treap) print() {
	if t.root == nil {
		return
	}

	queue := make([]*node, 0)

	queue = append(queue, t.root)

	for len(queue) != 0 {
		noOfNodes := len(queue)

		for i := 0; i < noOfNodes; i++ {
			current := queue[0]
			queue = queue[1:]

			fmt.Printf("%d(%d) ", current.peer.id, current.peer.currentCapacity)

			if current.left != nil {
				queue = append(queue, current.left)
			}

			if current.right != nil {
				queue = append(queue, current.right)
			}
		}

		fmt.Println()
	}

	return
}

func rightRotate(root *node) *node {
	left := root.left
	subTree := left.right

	left.right = root
	root.left = subTree

	return left
}

func leftRotate(root *node) *node {
	right := root.right
	subTree := right.left

	right.left = root
	root.right = subTree

	return right
}

func recursiveInsert(root *node, peer *peer) *node {
	if root == nil {
		return newNode(peer)
	}

	if peer.id <= root.peer.id {

		root.left = recursiveInsert(root.left, peer)

		if root.left != nil && root.left.peer.currentCapacity > root.peer.currentCapacity {
			root = rightRotate(root)
		}

		return root
	}

	root.right = recursiveInsert(root.right, peer)

	if root.right != nil && root.right.peer.currentCapacity > root.peer.currentCapacity {
		root = leftRotate(root)
	}

	return root
}

func recursiveDelete(root *node, id int) *node {
	if root == nil {
		return root
	}

	if id < root.peer.id {
		root.left = recursiveDelete(root.left, id)
		return root
	}

	if id > root.peer.id {
		root.right = recursiveDelete(root.right, id)
		return root
	}

	// no children
	if root.left == nil && root.right == nil {
		return nil
	}

	// having both children
	if root.left != nil && root.right != nil {

		if root.left.peer.currentCapacity < root.right.peer.currentCapacity {
			root = leftRotate(root)
			root.left = recursiveDelete(root.left, id)
			return root
		}

		root = rightRotate(root)
		root.right = recursiveDelete(root.right, id)
		return root

	}

	temp := root.left

	// having single child
	if root.right != nil {
		temp = root.right
	}

	root = temp

	return root
}
