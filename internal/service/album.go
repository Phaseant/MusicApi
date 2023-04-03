package service

import (
	"github.com/Phaseant/MusicAPI/entity"
	"github.com/Phaseant/MusicAPI/internal/repository"
)

type AlbumService struct {
	repo repository.Album
}

func newAlbumService(repo repository.Album) *AlbumService {
	return &AlbumService{repo: repo}
}

func (s *AlbumService) NewAlbum(album entity.Album) (string, error) {
	return s.repo.AddAlbum(album)
}

func (s *AlbumService) GetAlbum(id string) (entity.Album, error) {
	return s.repo.GetAlbum(id)
}

func (s *AlbumService) GetAllAlbums() ([]entity.Album, error) {
	return s.repo.GetAllAlbums()
}

func (s *AlbumService) DeleteAlbum(id string) bool {
	return s.repo.DeleteAlbum(id)
}
