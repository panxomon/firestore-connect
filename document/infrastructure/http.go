package document

import (
	"encoding/json"
	document "firestore-connect/document/domain"
	"net/http"
)

func (h *handler) GetDocument(idDocument string) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			p, err := h.service.Read(idDocument)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(&p)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
	return http.HandlerFunc(fn)
}

func (h *handler) CreateDocument(id string) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			w.Header().Set("Content-Type", "application/json")
			var d []document.Document
			err := json.NewDecoder(r.Body).Decode(&d)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			p, err := h.service.Create(id, d)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(&p)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}

	return http.HandlerFunc(fn)
}
