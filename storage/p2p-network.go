package storage

import (
	"fmt"

	"p2p-network-simulator/domain/interfaces"
)

type P2PNetwork struct {
	network []*tree
	treap   *treap
	ids     map[int]struct{}
}

func NewP2PNetwork() interfaces.P2PNetwork {
	return &P2PNetwork{
		network: make([]*tree, 0),
		treap:   newTreap(),
	}
}

func (network *P2PNetwork) Join() {
	peer1 := &peer{
		id:              1,
		maxCapacity:     2,
		currentCapacity: 0,
	}

	peer2 := &peer{
		id:              2,
		maxCapacity:     1,
		currentCapacity: 0,
	}

	peer3 := &peer{
		id:              3,
		maxCapacity:     3,
		currentCapacity: 0,
	}

	network.treap.insert(peer1)
	network.treap.insert(peer2)
	network.treap.insert(peer3)

	fmt.Println()

	network.treap.delete(peer1)
}

func (network *P2PNetwork) Leave() {
}

func (network *P2PNetwork) Trace() {
}
