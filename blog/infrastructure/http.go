package blog

import (
	"encoding/json"
	blog "firestore-connect/blog/domain"
	"net/http"
)

type Handler interface {
	Post(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	blogGateway blog.Gateway
}

func NewBlogHandler(gateway blog.Gateway) Handler {
	return &handler{blogGateway: gateway}
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("content-type", "application/json")

		var article = blog.Article{
			Id:     "chapter1",
			Title:  "Titulo",
			Text:   "Texto del blog",
			Date:   "20 de julio 2021",
			Author: "Pancho Montoya",
		}

		p, err := h.blogGateway.Create(article)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(&p)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		code := r.URL.Query()["key"]
		p, err := h.blogGateway.Read(code[0])

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(&p)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	p, err := h.blogGateway.GetAll()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(&p)
}
