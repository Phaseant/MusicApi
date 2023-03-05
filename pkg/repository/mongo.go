package repository

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	Username string
	Password string
}

const (
	DBName   = "MusicServer"
	UserCol  = "users"
	AlbumCol = "albums"
	AdminCol = "admins"
)

func InitMongo(cfg Config) (*mongo.Client, error) {

	credential := options.Credential{
		Username: cfg.Username,
		Password: cfg.Password,
	}

	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, err
	}

	log.Info("MongoDB is connected")
	return client, err
}
