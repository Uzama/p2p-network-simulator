package storage

import (
	"testing"

	"p2p-network-simulator/domain/entities"
)

var n = entities.Node{
	Id:       2,
	Capacity: 2,
}

var n2 = entities.Node{
	Id:       3,
	Capacity: 3,
}

var n3 = entities.Node{
	Id:       4,
	Capacity: 4,
}

var n4 = entities.Node{
	Id:       5,
	Capacity: 5,
}

func Test_newPeer(t *testing.T) {
	peer := newPeer(n)

	if peer.id != 2 || peer.maxCapacity != 2 || peer.currentCapacity != 2 || peer.parent != nil {
		t.Error("something wrong")
	}
}

func Test_setParent(t *testing.T) {
	peer := newPeer(n)

	parent := newPeer(n2)

	peer.setParent(parent)

	if peer.parent.id != 3 || peer.parent.maxCapacity != 3 || peer.parent.currentCapacity != 3 {
		t.Error("something wrong")
	}
}

func Test_addChild(t *testing.T) {
	peer := newPeer(n)

	child := newPeer(n2)

	err := peer.addChild(child)
	if err != nil {
		t.Error("error not expected")
	}

	if peer.currentCapacity != 1 || len(peer.children) != 1 {
		t.Error("something wrong")
	}

	peer.addChild(newPeer(n3))
	err = peer.addChild(newPeer(n4))

	if err.Error() != "not enough space to add" {
		t.Error("expecting error")
	}
}
