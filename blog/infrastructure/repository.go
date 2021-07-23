package blog

import (
	"context"
	blog "firestore-connect/blog/domain"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type firestoreRepository struct {
	collectionName string
}

func NewFirestoreRepository(collectionName string) blog.Repository {
	r := &firestoreRepository{collectionName: collectionName}
	return r
}

func (r *firestoreRepository) Create(a blog.Article) (*blog.Article, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("file.json")
	client, err := firestore.NewClient(ctx, "ProjectID", sa)
	if err != nil {
		log.Fatalf("fallo en crear un documento: %v", err)
		return nil, err
	}
	defer client.Close()
	_, err = client.Collection("blog").Doc(a.Id).Set(ctx, map[string]interface{}{
		"Title":  a.Title,
		"Text":   a.Text,
		"Author": a.Author,
		"Date":   a.Date,
	})
	if err != nil {
		log.Fatalf("no agrego ni una wea porque: %v", err)
		return nil, err
	}
	return &a, nil
}

func (r *firestoreRepository) Read(id string) (*blog.Article, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("file.json")
	client, err := firestore.NewClient(ctx, "ProjectID", sa)

	if err != nil {
		log.Fatalf("fallo en crear client: %v", err)
		return nil, err
	}
	defer client.Close()

	dataSnap, err := client.Collection(r.collectionName).Doc(id).Get(ctx)
	if err != nil {
		log.Fatalf("fallo en obtener collection: %v", err)
		return nil, err
	}
	var article blog.Article
	dataSnap.DataTo(&article)

	return &article, nil
}

func (r *firestoreRepository) GetAll() ([]blog.Article, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("file.json")
	client, err := firestore.NewClient(ctx, "ProjectID", sa)

	if err != nil {
		log.Fatalf("fallo en crear un documento: %v", err)
		return nil, err
	}
	defer client.Close()
	var articles []blog.Article
	iter := client.Collection(r.collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("error iterando...")
			return nil, err
		}
		a := blog.Article{
			Title:  doc.Data()["Title"].(string),
			Text:   doc.Data()["Text"].(string),
			Date:   doc.Data()["Date"].(string),
			Author: doc.Data()["Author"].(string),
		}
		articles = append(articles, a)
	}

	return articles, nil
}
