package tree

import (
	"strconv"
)

type Tree struct {
	root *Peer
}

func NewTree(peer *Peer) *Tree {
	return &Tree{
		root: peer,
	}
}

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

func (t *Tree) Encode() string {
	if t.root == nil {
		return ""
	}

	return recursiveEncode(t.root)
}

func recursiveEncode(root *Peer) string {
	if root == nil {
		return ""
	}

	current := root.MaxCapacity - root.CurrentCapacity
	capacity := "(" + strconv.Itoa(current) + "/" + strconv.Itoa(root.MaxCapacity) + ")"

	temp := strconv.Itoa(root.Id) + capacity
	temp += "[ "

	for _, child := range root.Children {
		temp += recursiveEncode(child)
	}

	temp += " ]"
	return temp
}
