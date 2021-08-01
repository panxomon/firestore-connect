package document

import (
	document "firestore-connect/document/domain"
	"fmt"
	"log"
	"testing"
)

const configFile = "config/ltgc.json"

func TestReadDocumentFromFirestore(t *testing.T) {

	r := NewFirestorerepository(configFile, "Prj1", "col2", 5)

	result, err := r.Read("doc3")
	if err != nil {
		t.Errorf("error creando cliente firestore...%d\n", err)
	}

	log.Println(result)

}

func TestCreateNewDocument(t *testing.T) {

	r := NewFirestorerepository("json.json", "Prj1", "col2", 25)

	var documents = []document.Document{
		{
			Key:   "author",
			Value: "Panxo",
			Type:  "string",
		},
		{
			Key:   "body",
			Value: "lorea el ipsum 2",
			Type:  "string",
		},
		{
			Key:   "date",
			Value: "25de julio 2021",
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

	doc, err := r.Create("nuevoDocumento", documents)
	if err != nil {
		t.Errorf("error creando cliente firestore...%d\n", err)
	}

	fmt.Println(doc)

}
