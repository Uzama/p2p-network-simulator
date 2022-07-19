package tree

import (
	"errors"

	"p2p-network-simulator/domain/entities"
)

type Peer struct {
	Id              int
	MaxCapacity     int
	CurrentCapacity int
	Parent          *Peer
	Children        []*Peer
}

func NewPeer(node entities.Node) *Peer {
	return &Peer{
		Id:              node.Id,
		MaxCapacity:     node.Capacity,
		CurrentCapacity: node.Capacity,
		Children:        make([]*Peer, 0),
	}
}

func (p *Peer) SetParent(parent *Peer) {
	p.Parent = parent
}

func (p *Peer) AddChild(child *Peer) error {
	if p.CurrentCapacity == 0 {
		return errors.New("not enough space to add")
	}

	p.Children = append(p.Children, child)
	p.CurrentCapacity -= 1

	child.SetParent(p)

	return nil
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

	p.CurrentCapacity += 1

	child.SetParent(nil)
}
