package interfaces

type P2PNetwork interface {
	Join()
	Leave()
	Trace()
}