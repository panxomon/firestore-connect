package server

import (
	domain "firestore-connect/document/domain"
	Infrastructure "firestore-connect/document/infrastructure"
	document "firestore-connect/document/infrastructure"
	"net/http"
)

const Project_id = "ltgc-99d6e"

type routerMux struct {
	*http.ServeMux
}

func MyServeMux() *routerMux {

	handler := newHandler()
	mux := http.NewServeMux()

	mux.Handle("/document/:Id[string]", GET(handler.GetDocument("str1")))
	mux.Handle("/document/create", POST(handler.CreateDocument("str2")))

	return &routerMux{mux}
}

func newHandler() document.Handler {
	repo := Infrastructure.NewFirestorerepository("json.json", Project_id, "col2", 30)
	gateway := domain.NewService(repo)
	handler := Infrastructure.NewHandler(gateway)

	return handler
}

func GET(next http.HandlerFunc) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Write([]byte("hola desde get"))
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Something bad happened!"))
		}
	}

	return http.HandlerFunc(fn)
}

func POST(next http.HandlerFunc) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			w.Write([]byte("hola desde POST "))
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Something bad happened!"))
		}
	}

	return http.HandlerFunc(fn)
}
