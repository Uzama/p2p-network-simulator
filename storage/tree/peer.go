package tree

import (
	"p2p-network-simulator/domain/entities"
)

// Peer: it is a actual node in the peer to peer network
// Peers can have at most MaxCapacity of children and one parent
type Peer struct {
	Id          int
	MaxCapacity int
	Capacity    int // free capacity
	Parent      *Peer
	Children    []*Peer
}

// NewPeer: creates new peer. Initially, Capacity is equal to MaxCapacity
func NewPeer(node entities.Node) *Peer {
	return &Peer{
		Id:          node.Id,
		MaxCapacity: node.Capacity,
		Capacity:    node.Capacity,
		Children:    make([]*Peer, 0),
	}
}

// SetParent: resets the parent with the given peer
func (p *Peer) SetParent(parent *Peer) {
	p.Parent = parent
}

// AddChild: adds the given peer into Children list.
// After adding the child, capacity decreases by one.
func (p *Peer) AddChild(child *Peer) {
	p.Children = append(p.Children, child)
	p.Capacity -= 1

	child.SetParent(p)
}

// RemoveChild: removes the given child from the children list
func (p *Peer) RemoveChild(child *Peer) {
	var children []*Peer
	var exists bool

	for _, c := range p.Children {
		if child.Id == c.Id {
			exists = true
			continue
		}

		children = append(children, c)
	}

	p.Children = children

	// if the given child not in the list, then there are no changes happen
	if !exists {
		return
	}

	// the parent capacity will increments by one
	p.Capacity += 1

	// child's parent removes from the parent and set it to nil
	child.SetParent(nil)
}
