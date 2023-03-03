package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Phaseant/MusicAPI/entity"
	"github.com/Phaseant/MusicAPI/pkg/repository"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

const (
	passwordSalt = "gekkfoidsfosdf33rksm..34*2@"
	TokenExp     = time.Hour * 24 * 30
	userClaim    = "userID"
)

var tokenSalt = []byte("kaphwroeh33sdf##4(#*$)f//der")

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

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = TokenExp
	claims[userClaim] = user.Id
	claims["issuedAt"] = time.Now()

	signedToken, err := token.SignedString(tokenSalt)
	if err != nil {
		return "", err
	}

	log.Info("Token for user with id: ", user.Id.Hex(), " generated")
	return signedToken, nil
}

func (s *AuthService) ParseToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return tokenSalt, nil
	})

	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	_, isExists := claims[userClaim]
	if ok && token.Valid && isExists {
		return fmt.Sprintf("%s", claims["userID"]), err
	} else {
		return "", errors.New("no user id in claims")
	}
}

func generateHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(passwordSalt)))
}
