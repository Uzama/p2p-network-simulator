package http

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"p2p-network-simulator/domain/usecases"
	"p2p-network-simulator/storage"

	"github.com/gorilla/mux"
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

// Join: controller for join the network
func (hdl handler) Join(w http.ResponseWriter, r *http.Request) {
	// decode request body
	node, err := decodeRequest(r)
	if err != nil {
		log.Printf("error:%s\n", err.Error())

		handleError(w, err, http.StatusBadRequest)
		return
	}

	err = hdl.usecase.Join(node)
	if err != nil {
		log.Printf("error:%s\n", err.Error())

		handleError(w, err, http.StatusUnprocessableEntity)
		return
	}

	log.Printf("trace:node %d<%d> join the network\n", node.Id, node.Capacity)
	handle(w, "successfully joined", node.Id, http.StatusCreated)
}

// Join: controller for leave the network
func (hdl handler) Leave(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// retrive id from the request
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("error:%s\n", err.Error())

		handleError(w, err, http.StatusBadRequest)
		return
	}

	if id < 1 {
		err = errors.New("id must be a positive integer")
		log.Printf("error:%s\n", err.Error())

		handleError(w, err, http.StatusBadRequest)
		return
	}

	err = hdl.usecase.Leave(id)
	if err != nil {
		log.Printf("error:%s\n", err.Error())

		handleError(w, err, http.StatusUnprocessableEntity)
		return
	}

	log.Printf("trace:node %d leave the network\n", id)
	handle(w, "successfully left", id, http.StatusAccepted)
}

// Join: controller for get trace of the network
func (hdl handler) Trace(w http.ResponseWriter, r *http.Request) {
	trace := hdl.usecase.Trace()

	log.Println("trace:network trace sent")
	handle(w, "trace received", trace, http.StatusOK)
}
