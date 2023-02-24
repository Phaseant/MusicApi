package service

import "github.com/Phaseant/MusicAPI/pkg/repository"

type Autorization interface {
}

type Album interface {
}

type Service struct {
	Autorization
	Album
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
