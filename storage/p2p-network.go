package storage

import (
	"fmt"
	"sync"

	"p2p-network-simulator/domain/entities"
	"p2p-network-simulator/domain/interfaces"
	"p2p-network-simulator/storage/treap"
	"p2p-network-simulator/storage/tree"
)

type P2PNetwork struct {
	topology []*tree.Tree
	treap    *treap.Treap
	ids      map[int]struct{}
	lock     sync.Mutex
}

func NewP2PNetwork() interfaces.P2PNetwork {
	return &P2PNetwork{
		topology: make([]*tree.Tree, 0),
		treap:    treap.NewTreap(),
		ids:      make(map[int]struct{}),
	}
}

func (network *P2PNetwork) Join(node entities.Node) error {
	network.lock.Lock()
	defer network.lock.Unlock()

	defer network.treap.Print()

	_, ok := network.ids[node.Id]
	if ok {
		return fmt.Errorf("id %d already reserved", node.Id)
	}

	peer := tree.NewPeer(node)

	network.ids[node.Id] = struct{}{}

	parentPeer := network.treap.MostCapacityPeer()

	if parentPeer == nil {
		tree := tree.NewTree(peer)

		network.topology = append(network.topology, tree)

		if peer.CurrentCapacity > 0 {
			network.treap.Insert(peer)
		}

		return nil
	}

	err := parentPeer.AddChild(peer)
	if err != nil {
		delete(network.ids, peer.Id)
		return err
	}

	network.treap.Delete(parentPeer.Id)

	if parentPeer.CurrentCapacity > 0 {
		network.treap.Insert(parentPeer)
	}

	if peer.CurrentCapacity > 0 {
		network.treap.Insert(peer)
	}

	return nil
}

func (network *P2PNetwork) Leave(id int) error {
	network.lock.Lock()
	defer network.lock.Unlock()

	defer network.treap.Print()

	var peer *tree.Peer

	for _, tree := range network.topology {
		peer = tree.Locate(id)

		if peer != nil {
			break
		}
	}

	if peer == nil {
		return fmt.Errorf("cannot locate id %d node", id)
	}

	// case-1: leaf node
	if len(peer.Children) == 0 {

		peer.Parent.RemoveChild(peer)

		network.treap.Delete(peer.Parent.Id)
		network.treap.Delete(peer.Id)

		network.treap.Insert(peer.Parent)

		delete(network.ids, peer.Id)
		return nil
	}

	// implement remaining cases

	return nil
}

func (network *P2PNetwork) Trace() []string {
	network.lock.Lock()
	defer network.lock.Unlock()

	var digram []string

	for _, tree := range network.topology {
		digram = append(digram, tree.Encode())
	}

	return digram
}
