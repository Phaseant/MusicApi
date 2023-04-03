package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Phaseant/MusicAPI/internal/handler"
	"github.com/Phaseant/MusicAPI/internal/repository"
	"github.com/Phaseant/MusicAPI/internal/server"
	"github.com/Phaseant/MusicAPI/internal/service"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/xlab/closer"
	"go.mongodb.org/mongo-driver/mongo"
)

var db *mongo.Client

func main() {
	closer.Bind(clearDB)

	setLogger()

	if err := initConfig(); err != nil {
		log.Fatalf("Error initiaizing config: %v", err)
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Unable to load env variables: %v", err)
	}

	db, err := repository.InitMongo(repository.Config{
		Username: os.Getenv("MONGODB_USERNAME"),
		Password: os.Getenv("MONGODB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	go func() {
		srv := new(server.Server)
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("Error while running server: %v", err)
		}

		closer.Close()
	}()

	closer.Hold()
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func clearDB() {
	fmt.Println("\nBye bye...")
	db.Disconnect(context.Background())
}

func setLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:03",
	})
}
