package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func initRouter() *mux.Router {
	r := mux.NewRouter()

	handler := newHandler()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	}).Methods(http.MethodGet)

	// define endpoints
	r.HandleFunc("/join", handler.Join).Methods(http.MethodPost)
	r.HandleFunc("/leave/{id}", handler.Leave).Methods(http.MethodDelete)
	r.HandleFunc("/trace", handler.Trace).Methods(http.MethodGet)

	return r
}
