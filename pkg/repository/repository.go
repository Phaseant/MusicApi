package repository

import (
	"github.com/Phaseant/MusicAPI/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type Autorization interface {
	NewUser(user entity.User) (string, error)
	GetUser(username, password string) (entity.User, error)
}

type Album interface {
	AddAlbum(album entity.Album) (string, error)
	GetAlbum(id string) (entity.Album, error)
	GetAllAlbums() ([]entity.Album, error)
	DeleteAlbum(id string) bool
}

type Repository struct {
	Autorization
	Album
}

func NewRepository(db *mongo.Client) *Repository {
	return &Repository{
		Autorization: NewAuthMongo(db),
		Album:        NewAlbumRepo(db),
	}
}
