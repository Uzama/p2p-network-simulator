package storage

import (
	"fmt"
	"sort"
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

	_, ok := network.ids[node.Id]
	if ok {
		return fmt.Errorf("id %d already reserved", node.Id)
	}

	peer := tree.NewPeer(node)

	err := network.addToNetwork(peer)
	if err != nil {
		return err
	}

	network.ids[node.Id] = struct{}{}

	return nil
}

func (network *P2PNetwork) Leave(id int) error {
	network.lock.Lock()
	defer network.lock.Unlock()

	var peer *tree.Peer
	var tree *tree.Tree

	for _, t := range network.topology {
		peer = t.Locate(id)
		tree = t

		if peer != nil {
			break
		}
	}

	if peer == nil {
		return fmt.Errorf("cannot locate id %d node", id)
	}

	return network.removeFromNetwork(peer, tree)
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

func (network *P2PNetwork) addToNetwork(peer *tree.Peer) error {
	parent := network.treap.MostCapacityPeer()

	if parent == nil {
		tree := tree.NewTree(peer)

		network.topology = append(network.topology, tree)

		if peer.CurrentCapacity > 0 {
			network.treap.Insert(peer)
		}

		return nil
	}

	err := parent.AddChild(peer)
	if err != nil {
		return err
	}

	network.treap.Delete(parent.Id)

	if parent.CurrentCapacity > 0 {
		network.treap.Insert(parent)
	}

	if peer.CurrentCapacity > 0 {
		network.treap.Insert(peer)
	}

	return nil
}

func (network *P2PNetwork) removeFromNetwork(peer *tree.Peer, tree *tree.Tree) error {
	parent := peer.Parent

	if parent != nil {
		parent.RemoveChild(peer)
	}

	delete(network.ids, peer.Id)

	network.treap.Delete(peer.Id)

	if len(peer.Children) == 0 {

		if parent == nil {
			network.removeTree(tree)
			return nil
		}

		network.treap.Delete(parent.Id)
		network.treap.Insert(parent)

		network.reArrange(parent, tree)

		return nil
	}

	if len(peer.Children) == 1 {
		nextChild := peer.Children[0]

		if parent == nil {
			tree.SetRoot(nextChild)
			nextChild.SetParent(nil)
			return nil
		}

		parent.AddChild(nextChild)

		network.reArrange(nextChild, tree)

		return nil
	}

	// sort the children according to most capacity
	sort.SliceStable(peer.Children, func(i, j int) bool {
		return peer.Children[i].CurrentCapacity > peer.Children[j].CurrentCapacity
	})

	// next child would be, child which has most capacity
	nextChild := peer.Children[0]

	if parent == nil {
		tree.SetRoot(nextChild)
		nextChild.SetParent(nil)
	} else {
		parent.AddChild(nextChild)
	}

	for _, child := range peer.Children[1:] {
		child.SetParent(nil)
		network.addToNetwork(child)
	}

	network.reArrange(nextChild, tree)

	return nil
}

func (network *P2PNetwork) reArrange(peer *tree.Peer, t *tree.Tree) {
	if peer == nil || peer.Parent == nil {
		return
	}

	canReArrange := peer.CurrentCapacity > (peer.Parent.CurrentCapacity + 1)

	if !canReArrange {
		return
	}

	parent := peer.Parent
	grandParent := parent.Parent

	parent.RemoveChild(peer)

	if grandParent == nil {
		t.SetRoot(peer)
		peer.SetParent(nil)
	} else {
		grandParent.RemoveChild(parent)
		grandParent.AddChild(peer)
	}

	peer.AddChild(parent)

	network.treap.Delete(parent.Id)
	network.treap.Delete(peer.Id)

	network.treap.Insert(parent)

	if peer.CurrentCapacity > 0 {
		network.treap.Insert(peer)
	}

	network.reArrange(peer, t)
	return
}

func (network *P2PNetwork) removeTree(t *tree.Tree) {
	topology := make([]*tree.Tree, 0)

	for _, tree := range network.topology {
		if tree.GetRoot().Id == t.GetRoot().Id {
			continue
		}

		topology = append(topology, tree)
	}

	network.topology = topology
}
