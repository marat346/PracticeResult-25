package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	cities "github.com/marat346/practiceResult-25-"
	"github.com/marat346/practiceResult-25-/tree/master/internal/repository"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {
	cities.InitLogging()

	if err := initConfig(); err != nil {
		log.Err(err).Msg("Configuration file not loaded")
	}

	log.Info().Msg("Loading data...")
	db, err := repository.NewDataBase(viper.GetString("csv.path"))
	if err != nil {
		log.Error().Msg(err.Error())
	}
	log.Info().Msg("Data loaded")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	log.Info().Msg("Starting server...")
	server := new(cities.Server)
	go func() {
		if err := server.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
			log.Err(err).Msg("Server is not running")
		}
	}()
	log.Info().Msg("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-quit

	log.Warn().Msg("Server shutting down...")

	log.Info().Msg("Saving CSV file...")
	if err := db.SaveCSV(viper.GetString("csv.path")); err != nil {
		log.Err(err).Msg("Configuration file not saved")
	}
	log.Info().Msg("СSV file saved")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Err(err).Msg("Server did not shut down correctly")
	}

	log.Info().Msg("Bye!")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
