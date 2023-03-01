package service

import (
	"github.com/Phaseant/MusicAPI/entity"
	"github.com/Phaseant/MusicAPI/pkg/repository"
)

type Autorization interface {
	NewUser(user entity.User) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (string, error)
}

type Album interface {
	NewAlbum(album entity.Album) (string, error)
	GetAlbum(id string) (entity.Album, error)
	GetAllAlbums() ([]entity.Album, error)
	DeleteAlbum(id string) bool
}

type Service struct {
	Autorization
	Album
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: newAuthService(repos.Autorization),
		Album:        newAlbumService(repos.Album),
	}
}
