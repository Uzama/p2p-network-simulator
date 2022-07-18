package storage

import (
	"testing"
)

/*
t1:
		7
	   / \
	  6   8
	     / \
		9  10
*/

var (
	t1 = newTree(p7)
	t2 = newTree(nil)
	t3 = &tree{root: p3}
	t4 = &tree{root: p4}
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
	testTable := []struct {
		name     string
		peer     *peer
		expected *tree
	}{
		{
			name:     "nil root",
			peer:     nil,
			expected: t2,
		},
		{
			name:     "happy case 1",
			peer:     p3,
			expected: t3,
		},
		{
			name:     "happy case 2",
			peer:     p4,
			expected: t4,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := newTree(testCase.peer)

			if result.root == nil && testCase.expected.root != nil {
				t.Errorf("expected %d , but got %v", testCase.expected.root.id, result.root)
			}

			if result.root != nil && testCase.expected.root == nil {
				t.Errorf("expected %v, but got %d", testCase.expected.root, result.root.id)
			}

			if result.root != nil && testCase.expected.root != nil && result.root.id != testCase.expected.root.id {
				t.Errorf("expected %d, but got %d", testCase.expected.root.id, result.root.id)
			}
		})
	}
}

func Test_locate(t *testing.T) {
	testTable := []struct {
		name     string
		tree     *tree
		id       int
		expected *peer
	}{
		{
			name:     "not exists id",
			tree:     t1,
			id:       2,
			expected: nil,
		},
		{
			name:     "happy case 1",
			tree:     t1,
			id:       9,
			expected: p9,
		},
		{
			name:     "happy case 2",
			tree:     t1,
			id:       6,
			expected: p6,
		},
		{
			name:     "nil tree",
			tree:     t2,
			id:       1,
			expected: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.tree.locate(testCase.id)

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

func Test_encode(t *testing.T) {
	testTable := []struct {
		name     string
		tree     *tree
		expected string
	}{
		{
			name:     "happy case 1",
			tree:     t1,
			expected: "7(2/2)[ 6(0/1)[  ]8(2/3)[ 9(0/4)[  ]10(0/5)[  ] ] ]",
		},
		{
			name:     "happy case 2",
			tree:     t3,
			expected: "3(0/3)[  ]",
		},
		{
			name:     "empty tree",
			tree:     t2,
			expected: "",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.tree.encode()

			if result != testCase.expected {
				t.Errorf("expected %s, but got %s", testCase.expected, result)
			}
		})
	}
}
