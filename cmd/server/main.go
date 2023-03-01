package main

import (
	"context"
	"os"

	"github.com/Phaseant/MusicAPI/entity"
	"github.com/Phaseant/MusicAPI/pkg/handler"
	"github.com/Phaseant/MusicAPI/pkg/repository"
	"github.com/Phaseant/MusicAPI/pkg/service"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:03",
	})

	if err := initConfig(); err != nil {
		log.Fatalf("error initiaizing config: %v", err)
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error unable to load env variables: %v", err)
	}
	db, err := repository.InitMongo(repository.Config{
		Username: viper.GetString("mongodb.username"),
		Password: os.Getenv("MONGODB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("Error unable to connect to database: %v", err)
	}

	defer func() {
		if err = db.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Error disconnecting to the database: %v", err)
		}
	}()

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(entity.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running server: %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func initLogger(out *os.File) {
	log.SetFormatter(&log.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-05-04 15:04:03",
	})
	log.SetOutput(out)
}
