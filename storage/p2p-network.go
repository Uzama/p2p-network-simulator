package storage

import (
	"fmt"
	"sync"

	"p2p-network-simulator/domain/entities"
	"p2p-network-simulator/domain/interfaces"
)

type P2PNetwork struct {
	topology []*tree
	treap    *treap
	ids      map[int]struct{}
	lock     sync.Mutex
}

func NewP2PNetwork() interfaces.P2PNetwork {
	return &P2PNetwork{
		topology: make([]*tree, 0),
		treap:    newTreap(),
		ids:      make(map[int]struct{}),
	}
}

func (network *P2PNetwork) Join(node entities.Node) error {
	network.lock.Lock()
	defer network.lock.Unlock()

	defer network.treap.print()

	_, ok := network.ids[node.Id]
	if ok {
		return fmt.Errorf("id %d already reserved", node.Id)
	}

	peer := newPeer(node)

	network.ids[node.Id] = struct{}{}

	parentPeer := network.treap.mostCapacityPeer()

	if parentPeer == nil {
		tree := newTree(peer)

		network.topology = append(network.topology, tree)

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
	network.lock.Lock()
	defer network.lock.Unlock()

	defer network.treap.print()

	var peer *peer

	for _, tree := range network.topology {
		peer = tree.locate(id)

		if peer != nil {
			break
		}
	}

	if peer == nil {
		return fmt.Errorf("cannot locate id %d node", id)
	}

	// case-1: leaf node
	if len(peer.children) == 0 {

		peer.parent.removeChild(peer)

		network.treap.delete(peer.parent.id)
		network.treap.delete(peer.id)

		network.treap.insert(peer.parent)

		delete(network.ids, peer.id)
		return nil
	}

	// implement remaining cases

	return nil
}

func (network *P2PNetwork) Trace() []string {
	network.lock.Lock()

	var topology []*tree

	for _, tree := range network.topology {
		topology = append(topology, tree.clone())
	}

	network.lock.Unlock()

	var digram []string

	for _, tree := range topology {
		digram = append(digram, tree.encode())
	}

	return digram
}
