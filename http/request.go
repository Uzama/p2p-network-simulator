package http

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"p2p-network-simulator/domain/entities"
)

type Node struct {
	Id       int `json:"id"`
	Capacity int `json:"capacity"`
}

func decodeRequest(r *http.Request) (entities.Node, error) {
	node := entities.Node{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return node, err
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, &node)
	if err != nil {
		return node, err
	}

	if node.Id < 1 {
		return node, errors.New("id must be a positive integer")
	}

	if node.Capacity < 1 {
		return node, errors.New("capacity must be a positive integer")
	}

	return node, nil
}
