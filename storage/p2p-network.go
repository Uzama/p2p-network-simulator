package storage

import (
	"fmt"

	"p2p-network-simulator/domain/entities"
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

func (network *P2PNetwork) Join(node entities.Node) error {
	_, ok := network.ids[node.Id]
	if ok {
		return fmt.Errorf("id %d already reserved", node.Id)
	}

	peer := &peer{
		id:              node.Id,
		maxCapacity:     node.Capacity,
		currentCapacity: node.Capacity,
		children:        make([]*peer, 0),
	}

	parentPeer := network.treap.mostCapacityPeer()

	if parentPeer == nil {
		tree := newTree(peer)

		network.network = append(network.network, tree)
		network.treap.insert(peer)

		return nil
	}

	network.treap.delete(parentPeer)

	parentPeer.children = append(parentPeer.children, peer)
	parentPeer.currentCapacity -= 1

	if parentPeer.currentCapacity > 0 {
		network.treap.insert(parentPeer)
	}

	network.treap.insert(peer)

	return nil
}

func (network *P2PNetwork) Leave() {
}

func (network *P2PNetwork) Trace() {
}
