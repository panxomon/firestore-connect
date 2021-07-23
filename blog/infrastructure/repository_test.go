package blog

import (
	blog "firestore-connect/blog/domain"
	"log"
	"testing"
)

func TestCreateArticle(t *testing.T) {
	r := NewFirestoreRepository("blog")

	var article = blog.Article{
		Id:     "chapter1",
		Title:  "Titulo",
		Text:   "Texto del blog",
		Date:   "20 de julio 2021",
		Author: "Pancho Montoya",
	}

	_, err := r.Create(article)

	if err != nil {
		t.Errorf("error ...%d\n", err)
	}
}

func TestReadArticleById(t *testing.T) {
	r := NewFirestoreRepository("blog")
	doc, err := r.Read("chapter1")
	if err != nil {
		t.Errorf("error ...%d\n", err)
	}

	log.Println(doc)
}

func TestGetAllArticles(t *testing.T) {
	r := NewFirestoreRepository("blog")
	articles, err := r.GetAll()
	if err != nil {
		t.Errorf("error ...%d\n", err)
	}

	log.Println(articles)
}
