package main

import "p2p-network-simulator/storage"

func main() {
	p := storage.NewP2PNetwork()

	p.Join()
}