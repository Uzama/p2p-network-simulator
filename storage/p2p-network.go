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

// P2PNetwork: it represts the actual p2p network
type P2PNetwork struct {
	// the network topology contains list of trees
	topology []*tree.Tree

	// we use treap to store peers which has free capacity (capacity > 0).
	// we can easily identify the peer which has the most free capacity (the root node of the treap)
	treap *treap.Treap

	// keeps track of joint peer's id in a set data structure to ensure ids are unique
	ids map[int]struct{}

	// using mutex to prevent from the concurrent accesses to the network
	lock sync.Mutex
}

// NewP2PNetwork: creates new p2p network
func NewP2PNetwork() interfaces.P2PNetwork {
	return &P2PNetwork{
		topology: make([]*tree.Tree, 0),
		treap:    treap.NewTreap(),
		ids:      make(map[int]struct{}),
	}
}

// Join: a new node joining the network
func (network *P2PNetwork) Join(node entities.Node) error {
	// using locks to prevent from concurrent access
	network.lock.Lock()
	defer network.lock.Unlock()

	// check whether the given node id is already reserved
	_, ok := network.ids[node.Id]
	if ok {
		return fmt.Errorf("id %d already reserved", node.Id)
	}

	// creating a new peer with given values
	peer := tree.NewPeer(node)

	// add to the network
	network.add(peer)

	// update the new node id
	network.ids[node.Id] = struct{}{}

	return nil
}

// Leave: a node leaving the network
func (network *P2PNetwork) Leave(id int) error {
	// using locks to prevent from concurrent access
	network.lock.Lock()
	defer network.lock.Unlock()

	// locate the peer and the tree for the given id
	var peer *tree.Peer
	var tree *tree.Tree

	for _, t := range network.topology {
		peer = t.Locate(id)
		tree = t

		if peer != nil {
			break
		}
	}

	// if the given id is not in the topology, then return an error
	if peer == nil {
		return fmt.Errorf("cannot locate id %d node", id)
	}

	// remove the peer from the network
	network.remove(peer, tree)

	return nil
}

// Trace: returns the current status of the network
func (network *P2PNetwork) Trace() []string {
	// using locks to prevent from concurrent access
	network.lock.Lock()
	defer network.lock.Unlock()

	var digram []string

	for _, tree := range network.topology {
		digram = append(digram, tree.Encode())
	}

	return digram
}

// add: adds the given peer to the network
func (network *P2PNetwork) add(peer *tree.Peer) {
	// get peer which has the most free capacity from the treap
	parent := network.treap.Get()

	// if there are no peers with free capacity, then add the given peer into a new tree
	if parent == nil {
		// new peer will become the root of the tree. so parent should be nil
		peer.SetParent(nil)

		// create a new tree with the given peer
		tree := tree.NewTree(peer)

		// add the new tree into the network topology
		network.topology = append(network.topology, tree)

		// if the given peer has free capacity, then insert it into the treap
		if peer.Capacity > 0 {
			network.treap.Insert(peer)
		}

		return
	}

	// add the given peer into the children list of the parent peer which has the most free capacity
	parent.AddChild(peer)

	// update the parent peer in the treap by delete and re insert it
	network.treap.Delete(parent.Id)

	// only insert peers into the treap, if they have free capacity
	if parent.Capacity > 0 {
		network.treap.Insert(parent)
	}

	if peer.Capacity > 0 {
		network.treap.Insert(peer)
	}
}

// remove: removes the given peer from the network
func (network *P2PNetwork) remove(peer *tree.Peer, tree *tree.Tree) {
	parent := peer.Parent

	// if the leaving peer is not the root, then remove the leaving peer from its parent
	if parent != nil {
		parent.RemoveChild(peer)
	}

	// delete the leaving peer from ids map and treap
	delete(network.ids, peer.Id)
	network.treap.Delete(peer.Id)

	// CASE A: removes a leaf peer
	if len(peer.Children) == 0 {

		// CASE A-1: the leaving peer is the root of the tree. need to delete the entire tree
		if parent == nil {
			network.removeTree(tree)
			return
		}

		// CASE A-2: leaving peer is not the root of the tree

		// update the parent peer in the treap
		network.treap.Delete(parent.Id)
		network.treap.Insert(parent)

		// reorder the parent in the tree
		network.reOrder(parent, tree)

		return
	}

	// CASE B: removes a peer which has exactly one child
	if len(peer.Children) == 1 {
		// next child would be the child of the leaving peer
		nextChild := peer.Children[0]

		// CASE B-1: the leaving peer is the root of the tree.
		// need to set new child as the root of the tree
		if parent == nil {
			tree.SetRoot(nextChild)
			nextChild.SetParent(nil)

			return
		}

		// CASE B-2: the leaving peer is not the root of the tree.

		// add the next child to the children list of the leaving peer's parent
		parent.AddChild(nextChild)

		// reorder the next child in the tree
		network.reOrder(nextChild, tree)

		return
	}

	// CASE C: removes a peer which has more than one child

	// sort the children according to capacity
	sort.SliceStable(peer.Children, func(i, j int) bool {
		return peer.Children[i].Capacity > peer.Children[j].Capacity
	})

	// next child would be the child which has most capacity
	nextChild := peer.Children[0]

	// if the leaving peer is the root, then next child become the root of the tree
	if parent == nil {
		tree.SetRoot(nextChild)
		nextChild.SetParent(nil)
	}

	// if the leaving peer is not the root,
	// then add the next child to children list of the parent peer of the leaving peer
	if parent != nil {
		parent.AddChild(nextChild)
	}

	// remaining children would be added to the network
	for _, child := range peer.Children[1:] {
		// delete child's tree peers from the treap
		// to prevent from adding the child to its own tree
		network.treap.DeepDelete(child)

		// add to the network
		network.add(child)

		// re insert the deleted child's tree peers
		network.treap.DeepInsert(child)
	}

	// reorder the next child in the tree
	network.reOrder(nextChild, tree)

	return
}

// reOrder: recursively reorders the given peer on the given tree
// to make sure that the tree's depth would become smaller.
// the peer moving upwards based on its free capacity and its parent free capacity
func (network *P2PNetwork) reOrder(peer *tree.Peer, tree *tree.Tree) {
	if peer == nil || peer.Parent == nil {
		return
	}

	// if the given peer has sufficient free capacity (parent peer free capacity +1),
	// then reorder the peer with its parent (make the parent peer as child of the given peer)
	if !(peer.Capacity > (peer.Parent.Capacity + 1)) {
		return
	}

	parent := peer.Parent
	grandParent := parent.Parent

	/*
		  1(1/1)
			|		reorder 10		         10(2/3)
		  10(1/3)  ---------------->          /   \
			|                             1(0/1) 12(0/2)
		  12(0/2)
	*/

	// remove the given peer from the children list of the parent
	parent.RemoveChild(peer)

	// if the parent is root of the tree, then the given peer become the root of the tree
	if grandParent == nil {
		tree.SetRoot(peer)
		peer.SetParent(nil)
	}

	// if the parent is not root of the tree,
	// then add the given peer into the children list of its grand parent
	if grandParent != nil {
		grandParent.RemoveChild(parent)
		grandParent.AddChild(peer)
	}

	// add the parent to the children list of the given peer
	peer.AddChild(parent)

	// update the peer and its parent in the treap
	network.treap.Delete(parent.Id)
	network.treap.Delete(peer.Id)

	network.treap.Insert(parent)

	if peer.Capacity > 0 {
		network.treap.Insert(peer)
	}

	// keep doing
	network.reOrder(peer, tree)
}

// removeTree: removes the given tree from the network topology
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
