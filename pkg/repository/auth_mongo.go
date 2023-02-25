package repository

import (
	"context"

	"github.com/Phaseant/MusicAPI/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongo struct {
	db *mongo.Client
}

func NewAuthMongo(db *mongo.Client) *AuthMongo {
	return &AuthMongo{db: db}
}

func (r *AuthMongo) NewUser(user entity.User) (string, error) {
	collection := r.db.Database(DBName).Collection(UserCol)
	user.Id = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return "", err
	}
	return user.Id.Hex(), nil
}
