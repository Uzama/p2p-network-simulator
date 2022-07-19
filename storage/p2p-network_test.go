package storage

import (
	"errors"
	"testing"

	"p2p-network-simulator/domain/entities"
)

var network = NewP2PNetwork()

var (
	n1  = entities.Node{Id: 1, Capacity: 1}
	n2  = entities.Node{Id: 2, Capacity: 1}
	n10 = entities.Node{Id: 10, Capacity: 1}
	n11 = entities.Node{Id: 11, Capacity: 2}
	n12 = entities.Node{Id: 12, Capacity: 0}
	n13 = entities.Node{Id: 13, Capacity: 0}

	n3 = entities.Node{Id: 3, Capacity: 3}
	n4 = entities.Node{Id: 4, Capacity: 0}
	n5 = entities.Node{Id: 5, Capacity: 1}
	n6 = entities.Node{Id: 6, Capacity: 0}
	n7 = entities.Node{Id: 7, Capacity: 5}
	n8 = entities.Node{Id: 8, Capacity: 1}
	n9 = entities.Node{Id: 9, Capacity: 1}
)

func TestJoin(t *testing.T) {
	testTable := []struct {
		name          string
		node          entities.Node
		expected      []string
		expectedError error
	}{
		{
			/*
				1
			*/
			name:          "add node 1",
			node:          n1,
			expected:      []string{"1(0/1)"},
			expectedError: nil,
		},
		{
			/*
				1
				|
				2
			*/
			name:          "add node 2",
			node:          n2,
			expected:      []string{"1(1/1)[ 2(0/1) ]"},
			expectedError: nil,
		},
		{
			name:          "add already exists node",
			node:          n2,
			expected:      []string{"1(1/1)[ 2(0/1) ]"},
			expectedError: errors.New("id 2 already reserved"),
		},
		{
			/*
					1
					|
					2
					|
				    10
			*/
			name:          "add node 10",
			node:          n10,
			expected:      []string{"1(1/1)[ 2(1/1)[ 10(0/1) ] ]"},
			expectedError: nil,
		},
		{
			/*
					1
					|
					2
					|
				    10
					|
					11
			*/
			name:          "add node 11",
			node:          n11,
			expected:      []string{"1(1/1)[ 2(1/1)[ 10(1/1)[ 11(0/2) ] ] ]"},
			expectedError: nil,
		},
		{
			/*
					1
					|
					2
					|
				    10
					|
					11
					|
					12
			*/
			name:          "add node 12",
			node:          n12,
			expected:      []string{"1(1/1)[ 2(1/1)[ 10(1/1)[ 11(1/2)[ 12(0/0) ] ] ] ]"},
			expectedError: nil,
		},
		{
			/*
					1
					|
					2
					|
				    10
					|
					11
				   /  \
				  12  13
			*/
			name:          "add node 13",
			node:          n13,
			expected:      []string{"1(1/1)[ 2(1/1)[ 10(1/1)[ 11(2/2)[ 12(0/0) 13(0/0) ] ] ] ]"},
			expectedError: nil,
		},
		{
			/*
					1				3
					|
					2
					|
				    10
					|
					11
				   /  \
				  12  13
			*/
			name: "add node 3",
			node: n3,
			expected: []string{
				"1(1/1)[ 2(1/1)[ 10(1/1)[ 11(2/2)[ 12(0/0) 13(0/0) ] ] ] ]",
				"3(0/3)",
			},
			expectedError: nil,
		},
		{
			/*
					1				3
					|				|
					2				4
					|
				    10
					|
					11
				   /  \
				  12  13
			*/
			name: "add node 4",
			node: n4,
			expected: []string{
				"1(1/1)[ 2(1/1)[ 10(1/1)[ 11(2/2)[ 12(0/0) 13(0/0) ] ] ] ]",
				"3(1/3)[ 4(0/0) ]",
			},
			expectedError: nil,
		},
		{
			/*
					1				 3
					|			   /  \
					2			  4    5
					|
				    10
					|
					11
				   /  \
				  12  13
			*/
			name: "add node 5",
			node: n5,
			expected: []string{
				"1(1/1)[ 2(1/1)[ 10(1/1)[ 11(2/2)[ 12(0/0) 13(0/0) ] ] ] ]",
				"3(2/3)[ 4(0/0) 5(0/1) ]",
			},
			expectedError: nil,
		},
		{
			/*
					1				  3
					|			    / | \
					2			   4  5  6
					|
				    10
					|
					11
				   /  \
				  12  13
			*/
			name: "add node 6",
			node: n6,
			expected: []string{
				"1(1/1)[ 2(1/1)[ 10(1/1)[ 11(2/2)[ 12(0/0) 13(0/0) ] ] ] ]",
				"3(3/3)[ 4(0/0) 5(0/1) 6(0/0) ]",
			},
			expectedError: nil,
		},
		{
			/*
					1				  3
					|			    / | \
					2			   4  5  6
					|				  |
				    10				  7
					|
					11
				   /  \
				  12  13
			*/
			name: "add node 7",
			node: n7,
			expected: []string{
				"1(1/1)[ 2(1/1)[ 10(1/1)[ 11(2/2)[ 12(0/0) 13(0/0) ] ] ] ]",
				"3(3/3)[ 4(0/0) 5(1/1)[ 7(0/5) ] 6(0/0) ]",
			},
			expectedError: nil,
		},
		{
			/*
					1				  3
					|			    / | \
					2			   4  5  6
					|				  |
				    10				  7
					|				  |
					11				  8
				   /  \
				  12  13
			*/
			name: "add node 8",
			node: n8,
			expected: []string{
				"1(1/1)[ 2(1/1)[ 10(1/1)[ 11(2/2)[ 12(0/0) 13(0/0) ] ] ] ]",
				"3(3/3)[ 4(0/0) 5(1/1)[ 7(1/5)[ 8(0/1) ] ] 6(0/0) ]",
			},
			expectedError: nil,
		},
		{
			/*
					1				  3
					|			    / | \
					2			   4  5  6
					|				  |
				    10				  7
					|				 / \
					11				8   9
				   /  \
				  12  13
			*/
			name: "add node 9",
			node: n9,
			expected: []string{
				"1(1/1)[ 2(1/1)[ 10(1/1)[ 11(2/2)[ 12(0/0) 13(0/0) ] ] ] ]",
				"3(3/3)[ 4(0/0) 5(1/1)[ 7(2/5)[ 8(0/1) 9(0/1) ] ] 6(0/0) ]",
			},
			expectedError: nil,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := network.Join(testCase.node)

			if result == nil && testCase.expectedError != nil {
				t.Errorf("expected %s, but got %v", testCase.expectedError.Error(), result)
			}

			if result != nil && testCase.expectedError == nil {
				t.Errorf("expected %v, but got %s", testCase.expectedError, result.Error())
			}

			if result != nil && testCase.expectedError != nil && result.Error() != testCase.expectedError.Error() {
				t.Errorf("expected %s, but got %s", testCase.expectedError.Error(), result.Error())
			}

			expected := make(map[string]struct{})

			for _, str := range testCase.expected {
				expected[str] = struct{}{}
			}

			got := network.Trace()

			for _, str := range got {

				_, ok := expected[str]
				if !ok {
					t.Errorf("%s is not expected", str)
				}

				delete(expected, str)
			}

			if len(expected) > 0 {
				t.Errorf("expected %v, but got %v", expected, got)
			}
		})
	}
}

func TestLeave(t *testing.T) {
	testTable := []struct {
		name          string
		id            int
		expected      []string
		expectedError error
	}{
		{
			name: "delete node 2",
			id:   2,
			expected: []string{
				"1(1/1)[ 10(1/1)[ 11(2/2)[ 12(0/0) 13(0/0) ] ] ]",
				"3(3/3)[ 4(0/0) 5(1/1)[ 7(2/5)[ 8(0/1) 9(0/1) ] ] 6(0/0) ]",
			},
			expectedError: nil,
			/*
					1				  3
					|			    / | \
					10			   4  5  6
					|				  |
				    11				  7
				   /  \				 / \
				  12  13			8   9

			*/
		},
		{
			name: "delete not exists node",
			id:   2,
			expected: []string{
				"1(1/1)[ 10(1/1)[ 11(2/2)[ 12(0/0) 13(0/0) ] ] ]",
				"3(3/3)[ 4(0/0) 5(1/1)[ 7(2/5)[ 8(0/1) 9(0/1) ] ] 6(0/0) ]",
			},
			expectedError: errors.New("cannot locate id 2 node"),
		},
		{
			name: "delete node 11",
			id:   11,
			expected: []string{
				"1(1/1)[ 10(1/1)[ 12(0/0) ] ]",
				"3(3/3)[ 4(0/0) 5(1/1)[ 7(3/5)[ 8(0/1) 9(0/1) 13(0/0) ] ] 6(0/0) ]",
			},
			expectedError: nil,
			/*
					1				  3
					|			    / | \
					10			   4  5  6
					|				  |
				    12				  7
				 					/ | \
				  				   8  9  13

			*/
		},
		{
			name: "delete node 10",
			id:   10,
			expected: []string{
				"1(1/1)[ 12(0/0) ]",
				"3(3/3)[ 4(0/0) 5(1/1)[ 7(3/5)[ 8(0/1) 9(0/1) 13(0/0) ] ] 6(0/0) ]",
			},
			expectedError: nil,
			/*
					1				  3
					|			    / | \
					12			   4  5  6
					 				  |
				    				  7
				 					/ | \
				  				   8  9  13

			*/
		},
		{
			name: "delete node 13",
			id:   13,
			expected: []string{
				"1(1/1)[ 12(0/0) ]",
				"7(4/5)[ 8(0/1) 9(0/1) 5(0/1) 3(2/3)[ 4(0/0) 6(0/0) ] ]",
			},
			expectedError: nil,
			/*
				1				   7
				|			    / / \ \
				12			   8 9  5  3
									  / \
									 4   6
			*/
		},
		{
			name: "delete node 7",
			id:   7,
			expected: []string{
				"1(1/1)[ 12(0/0) ]",
				"8(1/1)[ 3(3/3)[ 4(0/0) 6(0/0) 5(1/1)[ 9(0/1) ] ] ]",
			},
			expectedError: nil,
			/*
				1				   8
				|			   	   |
				12			   	   3
								 / | \
								4  6  5
								      |
									  9
			*/
		},
		{
			name: "delete node 1",
			id:   1,
			expected: []string{
				"12(0/0)",
				"8(1/1)[ 3(3/3)[ 4(0/0) 6(0/0) 5(1/1)[ 9(0/1) ] ] ]",
			},
			expectedError: nil,
			/*
				12				   8
							   	   |
							   	   3
								 / | \
								4  6  5
								      |
									  9
			*/
		},
		{
			name: "delete node 12",
			id:   12,
			expected: []string{
				"8(1/1)[ 3(3/3)[ 4(0/0) 6(0/0) 5(1/1)[ 9(0/1) ] ] ]",
			},
			expectedError: nil,
			/*
					   8
				   	   |
				   	   3
					 / | \
					4  6  5
					      |
						  9
			*/
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			result := network.Leave(testCase.id)

			if result == nil && testCase.expectedError != nil {
				t.Errorf("expected %s, but got %v", testCase.expectedError.Error(), result)
			}

			if result != nil && testCase.expectedError == nil {
				t.Errorf("expected %v, but got %s", testCase.expectedError, result.Error())
			}

			if result != nil && testCase.expectedError != nil && result.Error() != testCase.expectedError.Error() {
				t.Errorf("expected %s, but got %s", testCase.expectedError.Error(), result.Error())
			}

			expected := make(map[string]struct{})

			for _, str := range testCase.expected {
				expected[str] = struct{}{}
			}

			got := network.Trace()

			for _, str := range got {

				_, ok := expected[str]
				if !ok {
					t.Errorf("%s is not expected", str)
				}

				delete(expected, str)
			}

			if len(expected) > 0 {
				t.Errorf("expected %v, but got %v", expected, got)
			}
		})
	}
}
