package service

import (
	"github.com/Phaseant/MusicAPI/entity"
	"github.com/Phaseant/MusicAPI/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

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

type Admin interface {
	IsAdmin(id string) bool
	AddAdmin(admin entity.Admin) error
}

type Service struct {
	Autorization
	Album
	Admin
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: newAuthService(repos.Autorization),
		Album:        newAlbumService(repos.Album),
		Admin:        newAdminService(repos.Admin),
	}
}
