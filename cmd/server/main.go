package main

import (
	"log"

	"github.com/Phaseant/MusicAPI"
	"github.com/Phaseant/MusicAPI/pkg/handler"
	"github.com/Phaseant/MusicAPI/pkg/repository"
	"github.com/Phaseant/MusicAPI/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initiaizing config: %v", err)
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(MusicAPI.Server)

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running server: %v", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("../../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
