package tree

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

	p1  = NewPeer(n1)
	p2  = NewPeer(n2)
	p3  = NewPeer(n3)
	p4  = NewPeer(n5)
	p5  = NewPeer(n5)
	p6  = NewPeer(n6)
	p7  = NewPeer(n7)
	p8  = NewPeer(n8)
	p9  = NewPeer(n9)
	p10 = NewPeer(n10)
)

func TestNewPeer(t *testing.T) {
	testTable := []struct {
		name     string
		node     entities.Node
		expected *Peer
	}{
		{
			name:     "happy case 1",
			node:     n1,
			expected: p1,
		},
		{
			name:     "happy case 2",
			node:     n7,
			expected: p7,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := NewPeer(testCase.node)

			if result.Id != testCase.expected.Id {
				t.Errorf("expected %d, but got %d", testCase.expected.Id, result.Id)
			}
		})
	}
}

func TestSetParent(t *testing.T) {
	testTable := []struct {
		name   string
		child  *Peer
		parent *Peer
	}{
		{
			name:   "happy case",
			child:  p2,
			parent: p1,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.child.SetParent(testCase.parent)

			if testCase.child.Parent.Id != testCase.parent.Id {
				t.Errorf("expected %d, but got %d", testCase.parent.Id, testCase.child.Parent.Id)
			}
		})
	}
}

func TestAddChild(t *testing.T) {
	testTable := []struct {
		name   string
		child  *Peer
		parent *Peer
	}{
		{
			name:   "happy case 1",
			child:  p3,
			parent: p2,
		},
		{
			name:   "happy case 2",
			child:  p4,
			parent: p2,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.parent.AddChild(testCase.child)

			expectedCapacity := testCase.parent.MaxCapacity - len(testCase.parent.Children)

			if testCase.parent.Capacity != expectedCapacity {
				t.Errorf("expected %d, but got %d", expectedCapacity, testCase.parent.Capacity)
			}

			if testCase.child.Parent.Id != testCase.parent.Id {
				t.Errorf("expected %d, but got %d", testCase.parent.Id, testCase.child.Parent.Id)
			}
		})
	}
}

func TestRemoveChild(t *testing.T) {
	testTable := []struct {
		name   string
		child  *Peer
		parent *Peer
	}{
		{
			name:   "happy case",
			child:  p3,
			parent: p2,
		},
		{
			name:   "happy case2",
			child:  p4,
			parent: p2,
		},
		{
			name:   "not exists",
			child:  p5,
			parent: p2,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			currentCapacity := testCase.parent.Capacity

			testCase.parent.RemoveChild(testCase.child)

			expectedCapacity := testCase.parent.MaxCapacity - len(testCase.parent.Children)

			if testCase.parent.Capacity != expectedCapacity {
				t.Errorf("expected %d, but got %d", expectedCapacity, testCase.parent.Capacity)
			}

			if testCase.parent.Capacity != currentCapacity && testCase.child.Parent != nil {
				t.Errorf("expected nil, but got %d", testCase.child.Parent.Id)
			}
		})
	}
}
