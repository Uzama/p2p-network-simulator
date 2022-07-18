package storage

import (
	"testing"

	"p2p-network-simulator/domain/entities"
)

var (
	n1  = entities.Node{Id: 1, Capacity: 1}
	n2  = entities.Node{Id: 2, Capacity: 2}
	n3  = entities.Node{Id: 3, Capacity: 3}
	n4  = entities.Node{Id: 4, Capacity: 4}
	n5  = entities.Node{Id: 5, Capacity: 5}
	n6  = entities.Node{Id: 6, Capacity: 1}
	n7  = entities.Node{Id: 7, Capacity: 2}
	n8  = entities.Node{Id: 8, Capacity: 3}
	n9  = entities.Node{Id: 9, Capacity: 4}
	n10 = entities.Node{Id: 10, Capacity: 5}

	node0  = &node{}
	node1  = &node{peer: p1}
	node2  = &node{peer: p2}
	node3  = newNode(p3)
	node4  = newNode(p4)
	node5  = newNode(p5)
	node6  = newNode(p6)
	node7  = newNode(p7)
	node8  = newNode(p8)
	node9  = newNode(p9)
	node10 = newNode(p10)
)

func Test_newNode(t *testing.T) {
	testTable := []struct {
		name     string
		peer     *peer
		expected *node
	}{
		{
			name:     "nil peer",
			peer:     nil,
			expected: node0,
		},
		{
			name:     "happy case 1",
			peer:     p1,
			expected: node1,
		},
		{
			name:     "happy case 2",
			peer:     p2,
			expected: node2,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := newNode(testCase.peer)

			if result.peer == nil && testCase.expected.peer != nil {
				t.Errorf("expected %d , but got %v", testCase.expected.peer.id, result.peer)
			}

			if result.peer != nil && testCase.expected.peer == nil {
				t.Errorf("expected %v, but got %d", testCase.expected.peer, result.peer.id)
			}

			if result.peer != nil && testCase.expected.peer != nil && result.peer.id != testCase.expected.peer.id {
				t.Errorf("expected %d, but got %d", testCase.expected.peer.id, result.peer.id)
			}
		})
	}

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
	testTable := []struct {
		name     string
		node     *node
		expected *peer
	}{
		{
			name:     "happy case",
			node:     node1,
			expected: p1,
		},
		{
			name:     "nil root",
			node:     node0,
			expected: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.node.getPeer()

			if result == nil && testCase.expected != nil {
				t.Errorf("expected %d, but got %v", testCase.expected.id, result)
			}

			if result != nil && testCase.expected == nil {
				t.Errorf("expected %v, but got %d", testCase.expected, result.id)
			}

			if result != nil && testCase.expected != nil && result.id != testCase.expected.id {
				t.Errorf("expected %d, but got %d", testCase.expected.id, result.id)
			}
		})
	}
}
