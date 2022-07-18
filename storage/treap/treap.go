package treap

import (
	"fmt"

	"p2p-network-simulator/storage/tree"
)

type Treap struct {
	root *node
}

func NewTreap() *Treap {
	return &Treap{
		root: nil,
	}
}

func (t *Treap) Insert(peer *tree.Peer) {
	t.root = recursiveInsert(t.root, peer)
}

func (t *Treap) Delete(id int) {
	t.root = recursiveDelete(t.root, id)
}

func (t *Treap) MostCapacityPeer() *tree.Peer {
	if t.root == nil {
		return nil
	}

	return t.root.getPeer()
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

func recursiveInsert(root *node, peer *tree.Peer) *node {
	if root == nil {
		return newNode(peer)
	}

	if peer.Id <= root.peer.Id {

		root.left = recursiveInsert(root.left, peer)

		if root.left != nil && root.left.peer.CurrentCapacity > root.peer.CurrentCapacity {
			root = rightRotate(root)
		}

		return root
	}

	root.right = recursiveInsert(root.right, peer)

	if root.right != nil && root.right.peer.CurrentCapacity > root.peer.CurrentCapacity {
		root = leftRotate(root)
	}

	return root
}

func recursiveDelete(root *node, id int) *node {
	if root == nil {
		return root
	}

	if id < root.peer.Id {
		root.left = recursiveDelete(root.left, id)
		return root
	}

	if id > root.peer.Id {
		root.right = recursiveDelete(root.right, id)
		return root
	}

	// no children
	if root.left == nil && root.right == nil {
		return nil
	}

	// having both children
	if root.left != nil && root.right != nil {

		if root.left.peer.CurrentCapacity < root.right.peer.CurrentCapacity {
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

func (t *Treap) Print() {
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

			fmt.Printf("%d(%d) ", current.peer.Id, current.peer.CurrentCapacity)

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
