package blog

type Gateway interface {
	GetAll() ([]Article, error)
	Create(a Article) (*Article, error)
	Read(id string) (*Article, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Gateway {
	return &service{repo: r}
}

func (s *service) Create(a Article) (*Article, error) {
	return s.repo.Create(a)
}

func (s *service) Read(id string) (*Article, error) {
	return s.repo.Read(id)
}

func (s *service) GetAll() ([]Article, error) {
	return s.repo.GetAll()
}
