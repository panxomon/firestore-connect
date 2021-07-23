package document

import (
	"encoding/json"
	document "firestore-connect/document/domain"
	"net/http"
	"strings"
)

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")

		path := r.URL.Path
		m := strings.Split(path, "/")

		p, err := h.service.Read(m[2])
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(&p)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")

		var d = []document.Document{
			{
				Key:   "author",
				Value: "pancho",
			},
			{
				Key:   "body",
				Value: "lorea el ipsum",
			},
			{
				Key:   "date",
				Value: "16 de julio 2021",
			},
			{
				Key:   "title",
				Value: "titulo del articulo del blog del fin del mundo",
			},
		}

		p, err := h.service.Create(d)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(&p)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
