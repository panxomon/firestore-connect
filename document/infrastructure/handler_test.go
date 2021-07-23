package document

import (
	"log"
	"testing"
)

func TestHandler(t *testing.T) {
	r := NewFirestorerepository("ColecionNueva")
	h := NewHandler(r)
	if h == nil {
		t.Errorf("error creando handler..")
	}
	log.Println(h)
}
