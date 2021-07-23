package server

import (
	"log"
	"net/http"
)

type server struct {
	*http.Server
}

func Start(port string) {
	server := newServer(port)

	log.Println("Iniciando servidor en el puerto: " + port)

	log.Fatal(server.ListenAndServe())
}

func newServer(p string) *server {

	s := &http.Server{
		Addr: ":" + p,
	}

	return &server{s}
}
