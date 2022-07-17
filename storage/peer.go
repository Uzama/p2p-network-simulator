package storage

import (
	"errors"

	"p2p-network-simulator/domain/entities"
)

type peer struct {
	id              int
	maxCapacity     int
	currentCapacity int
	parent          *peer
	children        []*peer
}

func newPeer(node entities.Node) *peer {
	return &peer{
		id:              node.Id,
		maxCapacity:     node.Capacity,
		currentCapacity: node.Capacity,
		children:        make([]*peer, 0),
	}
}

func (p *peer) setParent(parent *peer) {
	p.parent = parent
}

func (p *peer) addChild(child *peer) error {
	if p.currentCapacity == 0 {
		return errors.New("not enough space to add")
	}

	p.children = append(p.children, child)
	p.currentCapacity -= 1

	return nil
}

func (p *peer) removeChild(child *peer) {
	var children []*peer

	for _, c := range p.children {
		if child.id == c.id {
			continue
		}

		children = append(children, c)
	}

	p.children = children
	p.currentCapacity += 1
}
