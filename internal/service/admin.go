package service

import (
	"github.com/Phaseant/MusicAPI/entity"
	"github.com/Phaseant/MusicAPI/internal/repository"
)

type AdminService struct {
	repo repository.Admin
}

func newAdminService(repo repository.Admin) *AdminService {
	return &AdminService{repo: repo}
}

func (s *AdminService) IsAdmin(id string) bool {
	return s.repo.IsAdmin(id)
}

func (s *AdminService) AddAdmin(admin entity.Admin) error {
	return s.repo.AddAdmin(admin)
}
