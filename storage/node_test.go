package storage

import (
	"testing"

	"p2p-network-simulator/domain/entities"
)

var n1 = entities.Node{
	Id:       1,
	Capacity: 1,
}

var n2 = entities.Node{
	Id:       2,
	Capacity: 2,
}

var n3 = entities.Node{
	Id:       3,
	Capacity: 3,
}

var n4 = entities.Node{
	Id:       4,
	Capacity: 4,
}

var n5 = entities.Node{
	Id:       5,
	Capacity: 5,
}

var n6 = entities.Node{
	Id:       6,
	Capacity: 1,
}

var n7 = entities.Node{
	Id:       7,
	Capacity: 2,
}

var n8 = entities.Node{
	Id:       8,
	Capacity: 3,
}

var n9 = entities.Node{
	Id:       9,
	Capacity: 4,
}

var n10 = entities.Node{
	Id:       10,
	Capacity: 5,
}

func Test_newNode(t *testing.T) {
	peer := newPeer(n1)

	node := newNode(peer)

	if node.peer.id != 1 || node.peer.maxCapacity != 1 {
		t.Error("something wrong")
	}

	if node.left != nil || node.right != nil {
		t.Error("left and right must be nil")
	}
}

func Test_getPeer(t *testing.T) {
	node := newNode(newPeer(n1))

	peer := node.getPeer()

	if peer.id != 1 || peer.maxCapacity != 1 {
		t.Error("invalid peer")
	}
}
