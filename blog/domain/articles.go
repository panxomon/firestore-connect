package blog

type Article struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	Date   string `json:"date"`
	Author string `json:"author"`
}

type Repository interface {
	Create(a Article) (*Article, error)
	Read(id string) (*Article, error)
	GetAll() ([]Article, error)
}
