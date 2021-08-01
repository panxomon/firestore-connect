package server

import (
	"log"
	"net/http"
)

type server struct {
	*http.Server
}

func Start(port string, mux *http.ServeMux) {
	server := newServer(port, mux)

	log.Println("Iniciando servidor en el puerto: " + port)

	log.Fatal(server.ListenAndServe())
}

func newServer(p string, mux *http.ServeMux) *server {

	s := &http.Server{
		Addr:    ":" + p,
		Handler: mux,
	}

	return &server{s}
}
