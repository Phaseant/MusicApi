package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/Phaseant/MusicAPI/entity"
	"github.com/Phaseant/MusicAPI/pkg/repository"
)

const salt = "gekkfoidsfosdf33rksm..34*2@"

type AuthService struct {
	repo repository.Autorization
}

func newAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) NewUser(user entity.User) (int, error) {
	user.Password = generateHash(user.Password)
	return s.repo.NewUser(user)
}

func generateHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
