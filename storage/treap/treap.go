package treap

import (
	"strconv"

	"p2p-network-simulator/storage/tree"
)

// Treap: binary search tree + heap (max heap)
// Binary search tree property: left sub tree keys are less than root + right sub tree keys
// Heap property: children priorities are less than the parent priority
// Use treap to keep track of peers which has the most free capacity
type Treap struct {
	root *node
}

// NewTreap: creates empty treap
func NewTreap() *Treap {
	return &Treap{
		root: nil,
	}
}

// Get: returns the peer which has the most free capacity (root node)
func (t *Treap) Get() *tree.Peer {
	if t.root == nil {
		return nil
	}

	return t.root.get()
}

// Insert: inserts the given peer into the treap.
// If the peer id already exists, then overwrites it
func (t *Treap) Insert(peer *tree.Peer) {
	t.root = recursiveInsert(t.root, peer)
}

// Delete: deletes the peer for given id from the treap.
// If peer is not exists in the treap, then there are no changes happen to the treap
func (t *Treap) Delete(id int) {
	t.root = recursiveDelete(t.root, id)
}

// DeepDelete: deletes each and every peer from the treap for the given tree
// Peer is a tree. Uses level order traversal to visits every node in the tree
func (t *Treap) DeepDelete(peer *tree.Peer) {
	queue := make([]*tree.Peer, 0)

	queue = append(queue, peer)

	for len(queue) != 0 {

		current := queue[0]
		queue = queue[1:]

		t.Delete(current.Id)

		if len(current.Children) > 0 {
			queue = append(queue, current.Children...)
		}
	}
}

// DeepInsert: inserts each and every peer into the treap for the given tree
// Peer is a tree. Uses level order traversal to visits every node in the tree
func (t *Treap) DeepInsert(peer *tree.Peer) {
	queue := make([]*tree.Peer, 0)

	queue = append(queue, peer)

	for len(queue) != 0 {

		current := queue[0]
		queue = queue[1:]

		// insert only if the peer has free capacity
		if current.Capacity > 0 {
			t.Insert(current)
		}

		if len(current.Children) > 0 {
			queue = append(queue, current.Children...)
		}
	}
}

/*
encode: encodes the treap as a string.
This will used in unit testing to validate the result

	node: <id:capacity>

		 4
		/ \
	   3   9
		  /
		 8

	(4:4)[ (3:2) (9:4)[ (8:3) ] ]
*/
func (t *Treap) encode() string {
	return recursiveEncode(t.root)
}

/*
rightRotate: does right rotation at given node to maintain the heap property in the treap

       root                      L
       / \     Right Rotate     / \
      L   R       ———>         X  root
     / \                          / \
    X   Y                        Y   R
*/
func rightRotate(root *node) *node {
	L := root.left
	Y := L.right

	L.right = root
	root.left = Y

	return L
}

/*
leftRotate: does left rotation at given node to maintain the heap property in the treap

     root                       R
     / \      Left Rotate      / \
    L   R        ———>       root  Y
       / \                   / \
      X   Y                 L   X

*/
func leftRotate(root *node) *node {
	R := root.right
	X := R.left

	R.left = root
	root.right = X

	return R
}

// recursiveInsert: recursively inserts the peer into the treap
func recursiveInsert(root *node, peer *tree.Peer) *node {
	if root == nil {
		return newNode(peer)
	}

	// if already exists, then overwrites it
	if peer.Id == root.peer.Id {
		root.peer.Capacity = peer.Capacity
		return root
	}

	// look for left sub tree
	if peer.Id < root.peer.Id {

		root.left = recursiveInsert(root.left, peer)

		// check whether the heap property effected
		if root.left != nil && root.left.peer.Capacity > root.peer.Capacity {
			root = rightRotate(root)
		}

		return root
	}

	// look for right sub tree
	root.right = recursiveInsert(root.right, peer)

	// check whether the heap property effected
	if root.right != nil && root.right.peer.Capacity > root.peer.Capacity {
		root = leftRotate(root)
	}

	return root
}

// recursiveDelete: recursively deletes the peer from the treap
func recursiveDelete(root *node, id int) *node {
	if root == nil {
		return root
	}

	// look for left sub tree
	if id < root.peer.Id {
		root.left = recursiveDelete(root.left, id)
		return root
	}

	// look for right sub tree
	if id > root.peer.Id {
		root.right = recursiveDelete(root.right, id)
		return root
	}

	// leaf node
	if root.left == nil && root.right == nil {
		return nil
	}

	// having both right & left child
	if root.left != nil && root.right != nil {

		// if the right child has more priority (capacity) than the left child,
		// then do left rotation around the root
		if root.left.peer.Capacity < root.right.peer.Capacity {
			root = leftRotate(root)
			root.left = recursiveDelete(root.left, id)
			return root
		}

		// otherwise do right rotation around the root
		root = rightRotate(root)
		root.right = recursiveDelete(root.right, id)
		return root

	}

	// having single child
	temp := root.left

	if root.right != nil {
		temp = root.right
	}

	root = temp

	return root
}

// recursiveEncode: recursively encodes the treap to a string
func recursiveEncode(root *node) string {
	if root == nil {
		return ""
	}

	temp := "(" + strconv.Itoa(root.peer.Id) + ":" + strconv.Itoa(root.peer.Capacity) + ")"

	if root.left == nil && root.right == nil {
		return temp
	}

	temp += "[ "

	temp += recursiveEncode(root.left)

	if root.left != nil && root.right != nil {
		temp += " "
	}

	temp += recursiveEncode(root.right)

	temp += " ]"
	return temp
}
