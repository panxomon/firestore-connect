package main

import (
	"fmt"
	"os"

	"firestore-connect/server"
)

const defaultPort = "3001"

func main() {
	fmt.Println("funciona")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// WIP: crear un router personalizado, es muy flaite la opcion del if r.Method ==

	// repo := blog.NewFirestoreRepository("blog")
	// gateway := blogdomain.NewService(repo)
	// handler := blog.NewBlogHandler(gateway)

	// http.HandleFunc("/articles/todes", handler.GetAll)
	// http.HandleFunc("/articles", handler.Get)
	// http.HandleFunc("/articles", handler.Post)

	server.Start(port)
}
