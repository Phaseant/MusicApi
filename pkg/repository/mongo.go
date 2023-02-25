package repository

import (
	"context"

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
)

func InitMongo(cfg Config) (*mongo.Client, error) {

	credential := options.Credential{
		Username: cfg.Username,
		Password: cfg.Password,
	}

	clientOpts := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(credential)

	client, err := mongo.Connect(context.TODO(), clientOpts) //check if connected
	if err != nil {
		return nil, err
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil { //check ping to verify that server is running
		return nil, err
	}
	return client, err
}
