package repository

type Autorization interface {
}

type Album interface {
}

type Repository struct {
	Autorization
	Album
}

func NewRepository() *Repository {
	return &Repository{}
}
