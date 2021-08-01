package document

import (
	"context"
	document "firestore-connect/document/domain"
	"fmt"
	"log"
	"reflect"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type firestoreRepository struct {
	collection string
	client     *firestore.Client
	timeout    time.Duration
}

func newClient(file string, projectID string, timeout int) (*firestore.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	sa := option.WithCredentialsFile(file)
	client, err := firestore.NewClient(ctx, projectID, sa)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewFirestorerepository(fileName string, projectID string, collectionName string, timeout int) document.Repository {
	client, err := newClient(fileName, projectID, 10)
	if err != nil {
		log.Println("error creando un cliente firestore %w", err)
	}

	return &firestoreRepository{
		collection: collectionName,
		client:     client,
		timeout:    time.Duration(timeout) * time.Second,
	}
}

func convertType(value interface{}) string {

	str := fmt.Sprintf("%v", value)
	return str
}

func (r *firestoreRepository) Create(id string, documents []document.Document) ([]document.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	data := make(map[string]interface{})
	for _, value := range documents {
		data[value.Key] = value.Value
	}

	d, err := r.client.Collection(r.collection).Doc(id).Set(ctx, data)
	if err != nil {
		log.Fatalf("Error agregando coleccion a cliente: %v", err)
		return nil, err
	}

	fmt.Println(d.UpdateTime)

	return documents, nil
}

func (r *firestoreRepository) Read(idCollection string) ([]document.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	dsnap, err := r.client.Collection(r.collection).Doc(idCollection).Get(ctx)
	if err != nil {
		return nil, err
	}
	var documents []document.Document
	m := dsnap.Data()

	for key, value := range m {
		var p document.Document
		p.Id = idCollection
		p.Key = key
		p.Value = convertType(value)
		p.Type = convertType(reflect.TypeOf(value))
		documents = append(documents, p)
	}

	return documents, nil
}

func (r *firestoreRepository) GetAll() ([]document.Document, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	var documents []document.Document

	iter := r.client.Collection(r.collection).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			log.Println("iterador done")
			break
		}
		if err != nil {
			log.Println("Failed to iterate: en el for err != nil pasado iterator.Done")
		}

		r := doc.Data()

		for key, value := range r {
			var p document.Document
			p.Key = key
			p.Value = value.(string)

			documents = append(documents, p)
		}

		fmt.Println(documents)
	}

	return documents, nil
}
