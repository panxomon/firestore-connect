package document

import (
	document "firestore-connect/document/domain"
	"net/http"
)

type handler struct {
	service document.Gateway
}

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
}

func NewHandler(gateway document.Gateway) Handler {
	return &handler{service: gateway}
}
