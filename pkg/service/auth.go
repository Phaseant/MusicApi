package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Phaseant/MusicAPI/entity"
	"github.com/Phaseant/MusicAPI/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	passwordSalt = "gekkfoidsfosdf33rksm..34*2@"
	tokenSalt    = "kaphwroeh33sdf##4(#*$)f//der"
	TokenExp     = time.Hour * 24 * 30
)

type CustomClaims struct {
	jwt.StandardClaims
	UserID string `json:"user_id"`
}

type AuthService struct {
	repo repository.Autorization
}

func newAuthService(repo repository.Autorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) NewUser(user entity.User) (string, error) {
	user.Password = generateHash(user.Password)
	return s.repo.NewUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generateHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExp).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id.Hex(),
	})
	return token.SignedString([]byte(tokenSalt))
}

func generateHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(passwordSalt)))
}
