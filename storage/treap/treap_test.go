package treap

import (
	"testing"

	"p2p-network-simulator/storage/tree"
)

var (
	treap  = NewTreap()
	treap2 = NewTreap()
)

func init() {
	treap2.Insert(p3)
	treap2.Insert(p4)
	treap2.Insert(p9)

	p3.AddChild(p9)
}

func TestInsert(t *testing.T) {
	testTable := []struct {
		name     string
		peer     *tree.Peer
		expected string
	}{
		{
			/*
				4
			*/
			name:     "add id 4",
			peer:     p4,
			expected: "(4:4)",
		},
		{
			/*
				4
				 \
				  8
			*/
			name:     "add id 8",
			peer:     p8,
			expected: "(4:4)[ (8:3) ]",
		},
		{
			/*
					4
				  /	 \
				 3	  8
			*/
			name:     "add id 3",
			peer:     p3,
			expected: "(4:4)[ (3:2) (8:3) ]",
		},
		{
			name:     "over write 3",
			peer:     p3,
			expected: "(4:4)[ (3:2) (8:3) ]",
		},
		{
			/*
					4
				  /	 \
				 3	  9
				     /
					8
			*/
			name:     "add id 9",
			peer:     p9,
			expected: "(4:4)[ (3:2) (9:4)[ (8:3) ] ]",
		},
		{
			/*
						5
					  /	 \
					 4	  9
				   /     /
				  3		8
			*/
			name:     "add id 5",
			peer:     p5,
			expected: "(5:5)[ (4:4)[ (3:2) ] (9:4)[ (8:3) ] ]",
		},
		{
			/*
						5
					  /	 \
					 4	  9
				   /     /
				  3		8
				       /
					  7
			*/
			name:     "add id 7",
			peer:     p7,
			expected: "(5:5)[ (4:4)[ (3:2) ] (9:4)[ (8:3)[ (7:2) ] ] ]",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			treap.Insert(testCase.peer)

			if treap.encode() != testCase.expected {
				t.Errorf("expected %s, but got %s", testCase.expected, treap.encode())
			}
		})
	}
}

func TestDelete(t *testing.T) {
	testTable := []struct {
		name     string
		id       int
		expected string
	}{
		{
			name:     "delete id 8",
			id:       8,
			expected: "(5:5)[ (4:4)[ (3:2) ] (9:4)[ (7:2) ] ]",
			/*
						5
					  /	 \
					 4	  9
				   /     /
				  3		7
			*/
		},
		{
			name:     "delete id 5",
			id:       5,
			expected: "(4:4)[ (3:2) (9:4)[ (7:2) ] ]",
			/*
						4
					  /	 \
					 3	  9
				         /
				    	7
			*/
		},
		{
			name:     "delete id not exists",
			id:       20,
			expected: "(4:4)[ (3:2) (9:4)[ (7:2) ] ]",
		},
		{
			name:     "delete id 7 (leaf node)",
			id:       7,
			expected: "(4:4)[ (3:2) (9:4) ]",
			/*
					4
				  /	 \
				 3	  9
			*/
		},
		{
			name:     "delete id 4",
			id:       4,
			expected: "(9:4)[ (3:2) ]",
			/*
					9
				  /
				 3
			*/
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			treap.Delete(testCase.id)

			if treap.encode() != testCase.expected {
				t.Errorf("expected %s, but got %s", testCase.expected, treap.encode())
			}
		})
	}
}

func TestMostCapacityPeer(t *testing.T) {
	testTable := []struct {
		name     string
		treap    *Treap
		expected *tree.Peer
	}{
		{
			name:     "happy case",
			treap:    treap,
			expected: p9,
		},
		{
			name:     "empty treap",
			treap:    NewTreap(),
			expected: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.treap.MostCapacityPeer()

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

func TestDeepDelete(t *testing.T) {
	testTable := []struct {
		name     string
		peer     *tree.Peer
		expected string
	}{
		{
			/*
					4
				  /  \
				 3    9
			*/
			name:     "delete id 3",
			peer:     p3,
			expected: "(4:4)",
			/*
				4

			*/
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			treap2.DeepDelete(testCase.peer)

			if treap2.encode() != testCase.expected {
				t.Errorf("expected %s, but got %s", testCase.expected, treap2.encode())
			}
		})
	}
}

func TestDeepInsert(t *testing.T) {
	testTable := []struct {
		name     string
		peer     *tree.Peer
		expected string
	}{
		{
			/*
				4

			*/
			name:     "insert id 3",
			peer:     p3,
			expected: "(4:4)[ (3:2) (9:4) ]",

			/*
					4
				  /  \
				 3    9
			*/
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			treap2.DeepInsert(testCase.peer)

			if treap2.encode() != testCase.expected {
				t.Errorf("expected %s, but got %s", testCase.expected, treap2.encode())
			}
		})
	}
}
