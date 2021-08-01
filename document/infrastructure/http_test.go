package document

import (
	document "firestore-connect/document/domain"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetDocument(t *testing.T) {

	r := NewFirestorerepository("json.json", "Prj1", "col2", 25)
	g := document.NewService(r)
	handler := NewHandler(g)

	req, err := http.NewRequest("GET", "/document", nil)

	if err != nil {
		t.Errorf("error creando  request para /document...%d\n", err)
	}

	rr := httptest.NewRecorder()
	h := http.HandlerFunc(handler.GetDocument("caca"))

	h.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
