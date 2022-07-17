package storage

import (
	"testing"
)

var (
	p1  = newPeer(n1)
	p2  = newPeer(n2)
	p3  = newPeer(n3)
	p4  = newPeer(n5)
	p5  = newPeer(n5)
	p6  = newPeer(n6)
	p7  = newPeer(n7)
	p8  = newPeer(n8)
	p9  = newPeer(n9)
	p10 = newPeer(n10)
)

func Test_newPeer(t *testing.T) {
	peer := newPeer(n2)

	if peer.id != 2 || peer.maxCapacity != 2 || peer.currentCapacity != 2 || peer.parent != nil {
		t.Error("something wrong")
	}
}

func Test_setParent(t *testing.T) {
	p2.setParent(p3)

	if p2.parent.id != 3 || p2.parent.maxCapacity != 3 || p2.parent.currentCapacity != 3 {
		t.Error("something wrong")
	}
}

func Test_addChild(t *testing.T) {
	err := p2.addChild(p3)
	if err != nil {
		t.Error("error not expected")
	}

	if p2.currentCapacity != 1 || len(p2.children) != 1 {
		t.Error("something wrong")
	}

	p2.addChild(p4)
	err = p2.addChild(p5)

	if err.Error() != "not enough space to add" {
		t.Error("expecting error")
	}
}
