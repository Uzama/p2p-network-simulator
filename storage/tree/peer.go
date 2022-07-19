package tree

import (
	"p2p-network-simulator/domain/entities"
)

type Peer struct {
	Id          int
	MaxCapacity int
	Capacity    int
	Parent      *Peer
	Children    []*Peer
}

func NewPeer(node entities.Node) *Peer {
	return &Peer{
		Id:          node.Id,
		MaxCapacity: node.Capacity,
		Capacity:    node.Capacity,
		Children:    make([]*Peer, 0),
	}
}

func (p *Peer) SetParent(parent *Peer) {
	p.Parent = parent
}

func (p *Peer) AddChild(child *Peer) {
	p.Children = append(p.Children, child)
	p.Capacity -= 1

	child.SetParent(p)
}

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

	if !exists {
		return
	}

	p.Capacity += 1

	child.SetParent(nil)
}
