package usecases

import (
	"p2p-network-simulator/domain/entities"
	"p2p-network-simulator/domain/interfaces"
)

type Simulator struct {
	network interfaces.P2PNetwork
}

func NewSimulator(network interfaces.P2PNetwork) Simulator {
	return Simulator{
		network: network,
	}
}

func (s Simulator) Join(node entities.Node) error {
	return s.network.Join(node)
}

func (s Simulator) Leave(id int) error {
	return s.network.Leave(id)
}

func (s Simulator) Trace() {
	s.network.Trace()
}
