package blog

import (
	"log"
	"net/http"
	"testing"
)

func TestPostNewArticle(t *testing.T) {

	req, err := http.NewRequest("GET", "/articles/l", nil)

	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	log.Println(req.URL)

}
