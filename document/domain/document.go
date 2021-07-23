package document

type Document struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Repository interface {
	Create(documents []Document) ([]Document, error)
	Read(idCollection string) ([]Document, error)
}
