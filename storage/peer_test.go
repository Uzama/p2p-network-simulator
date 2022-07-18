package storage

import (
	"errors"
	"testing"

	"p2p-network-simulator/domain/entities"
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
	testTable := []struct {
		name     string
		node     entities.Node
		expected *peer
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
			result := newPeer(testCase.node)

			if result.id != testCase.expected.id {
				t.Errorf("expected %d, but got %d", testCase.expected.id, result.id)
			}
		})
	}
}

func Test_setParent(t *testing.T) {
	testTable := []struct {
		name   string
		child  *peer
		parent *peer
	}{
		{
			name:   "happy case",
			child:  p2,
			parent: p1,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.child.setParent(testCase.parent)

			if testCase.child.parent.id != testCase.parent.id {
				t.Errorf("expected %d, but got %d", testCase.parent.id, testCase.child.parent.id)
			}
		})
	}
}

func Test_addChild(t *testing.T) {
	testTable := []struct {
		name     string
		child    *peer
		parent   *peer
		expected error
	}{
		{
			name:     "happy case 1",
			child:    p3,
			parent:   p2,
			expected: nil,
		},
		{
			name:     "happy case 2",
			child:    p4,
			parent:   p2,
			expected: nil,
		},
		{
			name:     "not enough space",
			child:    p5,
			parent:   p2,
			expected: errors.New("not enough space to add"),
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.parent.addChild(testCase.child)

			if result == nil && testCase.expected != nil {
				t.Errorf("expected %s, but got %v", testCase.expected.Error(), result)
			}

			if result != nil && testCase.expected == nil {
				t.Errorf("expected %v, but got %s", testCase.expected, result.Error())
			}

			if result != nil && testCase.expected != nil && result.Error() != testCase.expected.Error() {
				t.Errorf("expected %s, but got %s", testCase.expected.Error(), result.Error())
			}

			if result == nil && testCase.expected == nil {

				expectedCapacity := testCase.parent.maxCapacity - len(testCase.parent.children)

				if testCase.parent.currentCapacity != expectedCapacity {
					t.Errorf("expected %d, but got %d", expectedCapacity, testCase.parent.currentCapacity)
				}

			}
		})
	}
}
