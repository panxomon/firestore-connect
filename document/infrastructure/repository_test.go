package document

import (
	document "firestore-connect/document/domain"
	"fmt"
	"log"
	"testing"
)

func TestReadDocumentFromFirestore(t *testing.T) {

	r := NewFirestorerepository("col2")

	result, err := r.Read("doc3")
	if err != nil {
		t.Errorf("error creando cliente firestore...%d\n", err)
	}

	log.Println(result)

}

func TestCreateNewDocument(t *testing.T) {

	r := NewFirestorerepository("ColecionNueva")

	var documents = []document.Document{
		{
			Key:   "author",
			Value: "pancho",
			Type:  "string",
		},
		{
			Key:   "body",
			Value: "lorea el ipsum",
			Type:  "string",
		},
		{
			Key:   "date",
			Value: "16 de julio 2021",
			Type:  "string",
		},
		{
			Key:   "title",
			Value: "titulo del articulo del blog del fin del mundo",
			Type:  "string",
		},
		{
			Key:   "edad",
			Value: "39",
			Type:  "int",
		},
	}

	doc, err := r.Create(documents)
	if err != nil {
		t.Errorf("error creando cliente firestore...%d\n", err)
	}

	fmt.Println(doc)

}
