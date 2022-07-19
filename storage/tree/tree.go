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

func (t *Tree) GetRoot() *Peer {
	return t.root
}

func (t *Tree) SetRoot(peer *Peer) {
	t.root = peer
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
	return recursiveEncode(t.root)
}

func recursiveEncode(root *Peer) string {
	if root == nil {
		return ""
	}

	current := root.MaxCapacity - root.Capacity
	capacity := "(" + strconv.Itoa(current) + "/" + strconv.Itoa(root.MaxCapacity) + ")"

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
