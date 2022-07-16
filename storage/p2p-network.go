package storage

import "p2p-network-simulator/domain/interfaces"

type peer struct {
	id       string
	maxCapacity int
	currentCapacity int
	children []*peer
}

type tree struct {
	root   *peer
}

type P2PNetwork struct {
	network []*tree
	treap   *treap
}

func NewP2PNetwork() interfaces.P2PNetwork {
	return &P2PNetwork{
		network: make([]*tree, 0),
		treap: NewTreap(),
	}
}

func (network *P2PNetwork) Join() {

}

func (network *P2PNetwork) Leave() {
	
}

func (network *P2PNetwork) Trace() {
	
}
