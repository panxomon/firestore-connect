package document

import (
	"encoding/json"
	document "firestore-connect/document/domain"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDocument(t *testing.T) {

	req, err := http.NewRequest("GET", "/document/doc4", nil)

	if err != nil {
		t.Errorf("error creando  request para /document...%d\n", err)
	}

	r := NewFirestorerepository("col2")
	g := document.NewService(r)
	h := NewHandler(g)

	rec := httptest.NewRecorder()

	h.Get(rec, req)

	result := rec.Result()
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got: %d", http.StatusOK, result.StatusCode)
	}

	b, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	var got []*document.Document
	err = json.Unmarshal(b, &got)
	if err != nil {
		t.Fatalf("could not unmarshall response %v", err)
	}

	log.Println(got)

}
