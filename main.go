package main

import (
	"firestore-connect/server"
	"os"
)

const defaultPort = "3001"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	mux := server.MyServeMux()
	server.Start(port, mux.ServeMux)
}
