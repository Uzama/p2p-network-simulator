package storage

import (
	"fmt"
	"sync"

	"p2p-network-simulator/domain/entities"
	"p2p-network-simulator/domain/interfaces"
)

type P2PNetwork struct {
	network []*tree
	treap   *treap
	ids     map[int]struct{}
	lock    sync.Mutex
}

func NewP2PNetwork() interfaces.P2PNetwork {
	return &P2PNetwork{
		network: make([]*tree, 0),
		treap:   newTreap(),
	}
}

func (network *P2PNetwork) Join(node entities.Node) error {
	network.lock.Lock()
	defer network.lock.Unlock()

	_, ok := network.ids[node.Id]
	if ok {
		return fmt.Errorf("id %d already reserved", node.Id)
	}

	peer := newPeer(node)

	network.ids[node.Id] = struct{}{}

	parentPeer := network.treap.mostCapacityPeer()

	if parentPeer == nil {
		tree := newTree(peer)

		network.network = append(network.network, tree)

		if peer.currentCapacity > 0 {
			network.treap.insert(peer)
		}

		return nil
	}

	peer.setParent(parentPeer)

	network.treap.delete(parentPeer.id)

	parentPeer.addChild(peer)

	if parentPeer.currentCapacity > 0 {
		network.treap.insert(parentPeer)
	}

	if peer.currentCapacity > 0 {
		network.treap.insert(peer)
	}

	return nil
}

func (network *P2PNetwork) Leave(id int) error {
	return nil
}

func (network *P2PNetwork) Trace() {
}
