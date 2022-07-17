package storage

import (
	"strconv"

	"p2p-network-simulator/domain/entities"
)

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

func (t *tree) clone() *tree {
	if t.root == nil {
		return nil
	}

	return newTree(recursiveClone(t.root))
}

func recursiveClone(root *peer) *peer {
	if root == nil {
		return nil
	}

	var children []*peer

	for _, child := range root.children {
		children = append(children, recursiveClone(child))
	}

	peer := newPeer(entities.Node{
		Id:       root.id,
		Capacity: root.maxCapacity,
	})

	peer.setParent(root.parent)
	peer.setCurrentCapacity(root.currentCapacity)
	peer.setChildren(children)

	return peer
}

func (t *tree) encode() string {
	if t.root == nil {
		return ""
	}

	return recursiveEncode(t.root)
}

func recursiveEncode(root *peer) string {
	if root == nil {
		return ""
	}

	current := root.maxCapacity - root.currentCapacity
	capacity := "(" + strconv.Itoa(current) + "/" + strconv.Itoa(root.maxCapacity) + ")"

	temp := strconv.Itoa(root.id) + capacity
	temp += "[ "

	for _, child := range root.children {
		temp += recursiveEncode(child)
	}

	temp += " ]"
	return temp
}
