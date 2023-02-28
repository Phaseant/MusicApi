package service

import (
	"crypto/sha1"
	"errors"
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

type customClaims struct {
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExp).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id.Hex(),
	})
	return token.SignedString([]byte(tokenSalt))
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token signing method is not valid")
		}

		return []byte(tokenSalt), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return "", errors.New("token claims is not valid")
	}

	return claims.UserID, nil
}

func generateHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(passwordSalt)))
}
