package http

import (
	"net/http"

	"p2p-network-simulator/domain/usecases"
	"p2p-network-simulator/storage"
)

type handler struct {
	usecase usecases.Simulator
}

func newHandler() handler {
	network := storage.NewP2PNetwork()

	return handler{
		usecase: usecases.NewSimulator(network),
	}
}

func (hdl handler) Join(w http.ResponseWriter, r *http.Request) {
}

func (hdl handler) Leave(w http.ResponseWriter, r *http.Request) {
}

func (hdl handler) Trace(w http.ResponseWriter, r *http.Request) {
}
