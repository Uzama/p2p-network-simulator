package tree

import (
	"strconv"
)

// Tree: it is represents a sub network. It's connected with peers and form a n-ary tree
// Here, n in n-ary tree is equal to max capacity of each peer
type Tree struct {
	root *Peer
}

// NewTree: creates empty tree
func NewTree(peer *Peer) *Tree {
	return &Tree{
		root: peer,
	}
}

// GetRoot: returns the root
func (t *Tree) GetRoot() *Peer {
	return t.root
}

// SetRoot: resets the root with the given peer
func (t *Tree) SetRoot(peer *Peer) {
	t.root = peer
}

// Locate: returns the peer for given id. if id is not in the tree, then returns nil
// Uses level order traversal to visits every node in the tree
func (t *Tree) Locate(id int) *Peer {
	if t.root == nil {
		return nil
	}

	queue := make([]*Peer, 0)

	queue = append(queue, t.root)

	for len(queue) != 0 {

		current := queue[0]
		queue = queue[1:]

		if current.Id == id {
			return current
		}

		if len(current.Children) > 0 {
			queue = append(queue, current.Children...)
		}
	}

	return nil
}

/*
Encode: encodes the tree as a string.

	node: id ( #child / max capacity)

		7
	   / \
	  6   8
	     / \
		9  10

	7(2/2)[ 6(0/1) 8(2/3)[ 9(0/4) 10(0/5) ] ]
*/
func (t *Tree) Encode() string {
	return recursiveEncode(t.root)
}

// recursiveEncode: recursively encodes the tree to a string
func recursiveEncode(root *Peer) string {
	if root == nil {
		return ""
	}

	children := len(root.Children) // no of children
	capacity := "(" + strconv.Itoa(children) + "/" + strconv.Itoa(root.MaxCapacity) + ")"

	temp := strconv.Itoa(root.Id) + capacity

	if len(root.Children) == 0 {
		return temp
	}

	temp += "[ "

	for index, child := range root.Children {
		temp += recursiveEncode(child)

		if index != len(root.Children)-1 {
			temp += " "
		}
	}

	temp += " ]"
	return temp
}
