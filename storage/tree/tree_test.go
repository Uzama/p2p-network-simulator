package tree

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
	t1 = NewTree(p7)
	t2 = NewTree(nil)
	t3 = &Tree{root: p3}
	t4 = &Tree{root: p4}
)

func init() {
	t1.root.AddChild(p6)
	p6.SetParent(t1.root)

	t1.root.AddChild(p8)
	p8.SetParent(t1.root)

	p8.AddChild(p9)
	p9.SetParent(p8)

	p8.AddChild(p10)
	p10.SetParent(p8)
}

func TestNewTree(t *testing.T) {
	testTable := []struct {
		name     string
		peer     *Peer
		expected *Tree
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
			result := NewTree(testCase.peer)

			if result.root == nil && testCase.expected.root != nil {
				t.Errorf("expected %d , but got %v", testCase.expected.root.Id, result.root)
			}

			if result.root != nil && testCase.expected.root == nil {
				t.Errorf("expected %v, but got %d", testCase.expected.root, result.root.Id)
			}

			if result.root != nil && testCase.expected.root != nil && result.root.Id != testCase.expected.root.Id {
				t.Errorf("expected %d, but got %d", testCase.expected.root.Id, result.root.Id)
			}
		})
	}
}

func TestGetRoot(t *testing.T) {
	testTable := []struct {
		name     string
		tree     *Tree
		expected *Peer
	}{
		{
			name:     "nil root",
			tree:     t2,
			expected: nil,
		},
		{
			name:     "happy case",
			tree:     t1,
			expected: p7,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.tree.GetRoot()

			if result == nil && testCase.expected != nil {
				t.Errorf("expected %d , but got %v", testCase.expected.Id, result)
			}

			if result != nil && testCase.expected == nil {
				t.Errorf("expected %v, but got %d", testCase.expected, result.Id)
			}

			if result != nil && testCase.expected != nil && result.Id != testCase.expected.Id {
				t.Errorf("expected %d, but got %d", testCase.expected.Id, result.Id)
			}
		})
	}
}

func TestSetRoot(t *testing.T) {
	testTable := []struct {
		name     string
		peer     *Peer
		tree     *Tree
		expected *Tree
	}{
		{
			name:     "happy case",
			peer:     p3,
			tree:     t2,
			expected: t3,
		},
		{
			name:     "set nil",
			peer:     nil,
			tree:     t2,
			expected: t2,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.tree.SetRoot(testCase.peer)

			if testCase.tree.root == nil && testCase.expected.root != nil {
				t.Errorf("expected %d , but got %v", testCase.expected.root.Id, testCase.tree.root)
			}

			if testCase.tree.root != nil && testCase.expected.root == nil {
				t.Errorf("expected %v, but got %d", testCase.expected.root, testCase.tree.root.Id)
			}

			if testCase.tree.root != nil && testCase.expected.root != nil && testCase.tree.root.Id != testCase.expected.root.Id {
				t.Errorf("expected %d, but got %d", testCase.expected.root.Id, testCase.tree.root.Id)
			}
		})
	}
}

func TestLocate(t *testing.T) {
	testTable := []struct {
		name     string
		tree     *Tree
		id       int
		expected *Peer
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
			result := testCase.tree.Locate(testCase.id)

			if result == nil && testCase.expected != nil {
				t.Errorf("expected %d, but got %v", testCase.expected.Id, result)
			}

			if result != nil && testCase.expected == nil {
				t.Errorf("expected %v, but got %d", testCase.expected, result.Id)
			}

			if result != nil && testCase.expected != nil && result.Id != testCase.expected.Id {
				t.Errorf("expected %d, but got %d", testCase.expected.Id, result.Id)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	testTable := []struct {
		name     string
		tree     *Tree
		expected string
	}{
		{
			name:     "happy case 1",
			tree:     t1,
			expected: "7(2/2)[ 6(0/1) 8(2/3)[ 9(0/4) 10(0/5) ] ]",
		},
		{
			name:     "happy case 2",
			tree:     t3,
			expected: "3(0/3)",
		},
		{
			name:     "empty tree",
			tree:     t2,
			expected: "",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.tree.Encode()

			if result != testCase.expected {
				t.Errorf("expected %s, but got %s", testCase.expected, result)
			}
		})
	}
}
