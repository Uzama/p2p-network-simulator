package storage

import (
	"testing"
)

/*
		7
	   / \
	  6   8
	     / \
		9  10
*/

var (
	t1 = newTree(p7)
	t2 = newTree(nil)
)

func init() {
	t1.root.addChild(p6)
	p6.setParent(t1.root)

	t1.root.addChild(p8)
	p8.setParent(t1.root)

	p8.addChild(p9)
	p9.setParent(p8)

	p8.addChild(p10)
	p10.setParent(p8)
}

func Test_newTree(t *testing.T) {
	peer := newPeer(n1)

	tree := newTree(peer)

	if tree.root == nil {
		t.Error("somthing wrong")
	}
}

func Test_locate(t *testing.T) {
	peer := t1.locate(p2.id)

	if peer != nil {
		t.Error("expected nil")
	}

	peer = t1.locate(p9.id)

	if peer.id != 9 || peer.maxCapacity != 4 {
		t.Error("something wrong")
	}

	peer = t2.locate(p1.id)

	if peer != nil {
		t.Error("expected nil")
	}
}

func Test_encode(t *testing.T) {
	str := t1.encode()

	if str != "7(2/2)[ 6(0/1)[  ]8(2/3)[ 9(0/4)[  ]10(0/5)[  ] ] ]" {
		t.Error("un expected string")
	}

	str = t2.encode()

	if str != "" {
		t.Error("expected empty string")
	}
}
