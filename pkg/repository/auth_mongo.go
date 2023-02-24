package repository

import (
	"github.com/Phaseant/MusicAPI/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongo struct {
	db *mongo.Client
}

func NewAuthMongo(db *mongo.Client) *AuthMongo {
	return &AuthMongo{db: db}
}

func (r *AuthMongo) NewUser(user entity.User) (int, error) {
	return 0, nil
}
