package storage

import (
	"testing"

	"p2p-network-simulator/domain/entities"
)

var p = newPeer(entities.Node{
	Id:       1,
	Capacity: 1,
})

func Test_newNode(t *testing.T) {
	node := newNode(p)

	if node.peer.id != 1 || node.peer.maxCapacity != 1 {
		t.Error("something wrong")
	}

	if node.left != nil || node.right != nil {
		t.Error("left and right must be nil")
	}
}

func Test_getPeer(t *testing.T) {
	node := newNode(p)

	peer := node.getPeer()

	if peer.id != 1 || peer.maxCapacity != 1 {
		t.Error("invalid peer")
	}
}
