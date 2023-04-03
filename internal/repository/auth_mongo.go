package repository

import (
	"context"

	"github.com/Phaseant/MusicAPI/entity"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
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
	_, err := collection.InsertOne(context.Background(), user)

	if err != nil {
		return "", err
	}
	log.Info("Created new user with id: ", user.Id.Hex())
	return user.Id.Hex(), nil
}

func (r *AuthMongo) GetUser(username, password string) (entity.User, error) {
	collection := r.db.Database(DBName).Collection(UserCol)
	var user entity.User

	filter := bson.M{"$and": []interface{}{
		bson.D{{Key: "username", Value: username}},
		bson.D{{Key: "password", Value: password}},
	}}

	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil

}
