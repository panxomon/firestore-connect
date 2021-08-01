package document

type Gateway interface {
	Create(id string, d []Document) ([]Document, error)
	Read(idCollection string) ([]Document, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Gateway {
	return &service{repo: r}
}

func (s *service) Create(id string, d []Document) ([]Document, error) {
	return s.repo.Create(id, d)
}

func (s *service) Read(idCollection string) ([]Document, error) {
	return s.repo.Read(idCollection)
}
