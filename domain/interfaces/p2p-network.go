package interfaces

import "p2p-network-simulator/domain/entities"

type P2PNetwork interface {
	Join(node entities.Node) error
	Leave(id int) error
	Trace()
}
