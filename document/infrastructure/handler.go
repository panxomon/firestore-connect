package document

import (
	document "firestore-connect/document/domain"
	"net/http"
)

type handler struct {
	service document.Gateway
}

type Handler interface {
	GetDocument(id string) http.HandlerFunc
	CreateDocument(id string) http.HandlerFunc
}

func NewHandler(gateway document.Gateway) Handler {
	return &handler{service: gateway}
}
