package document

import (
	"log"
	"testing"
)

func TestHandler(t *testing.T) {
	r := NewFirestorerepository("json.json", "ltgc-99d6e", "col3", 10)
	h := NewHandler(r)
	if h == nil {
		t.Errorf("error creando handler..")
	}
	log.Println(h)
}
