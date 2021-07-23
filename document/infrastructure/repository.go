package document

import (
	"context"
	document "firestore-connect/document/domain"
	"fmt"
	"log"
	"reflect"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type firestoreRepository struct {
	collection string
}

func NewFirestorerepository(collectionName string) document.Repository {
	return &firestoreRepository{
		collection: collectionName,
	}

}

func convertType(value interface{}) string {

	str := fmt.Sprintf("%v", value)
	return str
}

func (r *firestoreRepository) Create(documents []document.Document) ([]document.Document, error) {

	ctx := context.Background()
	sa := option.WithCredentialsFile("file.json")

	client, err := firestore.NewClient(ctx, "ProjectID", sa)

	if err != nil {
		log.Fatalf("fallo en crear un documento: %v", err)
		return nil, err
	}

	data := make(map[string]interface{})
	for _, value := range documents {
		data[value.Key] = value.Value
	}
	d, err := client.Collection(r.collection).Doc("nuevo").Set(ctx, data)
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
		return nil, err
	}

	fmt.Println(d.UpdateTime)

	return documents, nil

}

func (r *firestoreRepository) Read(idCollection string) ([]document.Document, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("file.json")

	client, err := firestore.NewClient(ctx, "ProjectID", sa)
	if err != nil {
		log.Println("creando app.Firestore(ctx)")
		log.Fatalln(err)
	}

	var documents []document.Document
	defer client.Close()
	dsnap, err := client.Collection(r.collection).Doc(idCollection).Get(ctx)
	if err != nil {
		return nil, err
	}

	m := dsnap.Data()

	for key, value := range m {
		var p document.Document
		p.Key = key
		p.Value = convertType(value)
		p.Type = convertType(reflect.TypeOf(value))
		documents = append(documents, p)
	}

	return documents, nil
}

func (r *firestoreRepository) GetAll() ([]document.Document, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("file.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Println("creando new app (ctx, nil, sa)")
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Println("creando app.Firestore(ctx)")
		log.Fatalln(err)
	}

	var documents []document.Document

	iter := client.Collection(r.collection).Documents(ctx)

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
		defer client.Close()
	}

	return documents, nil
}
