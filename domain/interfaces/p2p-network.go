package interfaces

import "p2p-network-simulator/domain/entities"

type P2PNetwork interface {
	Join(peer entities.Node) error
	Leave()
	Trace()
}
