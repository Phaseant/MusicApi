package repository

import (
	"github.com/Phaseant/MusicAPI/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type Autorization interface {
	NewUser(user entity.User) (string, error)
	GetUser(username, password string) (entity.User, error)
}

type Album interface { //TODO
	AddAlbum(album entity.Album) (int, error)
	DeleteAlbum(album entity.Album) (int, error)
	GetAlbum(album entity.Album) (int, error)
}

type Repository struct {
	Autorization
	Album
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Autorization: NewAuthMongo(db),
	}
}
