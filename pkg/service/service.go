package service

import (
	"github.com/Phaseant/MusicAPI/entity"
	"github.com/Phaseant/MusicAPI/pkg/repository"
)

type Autorization interface {
	NewUser(user entity.User) (int, error)
}

type Album interface {
}

type Service struct {
	Autorization
	Album
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: newAuthService(repos.Autorization),
	}
}
