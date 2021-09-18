package main

import (
	"github.com/joho/godotenv"
	moneytracker "github.com/len3fun/money-tracker"
	"github.com/len3fun/money-tracker/pkg/handler"
	"github.com/len3fun/money-tracker/pkg/logger"
	"github.com/len3fun/money-tracker/pkg/repository"
	"github.com/len3fun/money-tracker/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logger.InitLogger()
	if err := initConfig(); err != nil {
		logger.Errorf("error initializing config: %s", err.Error())
		return
	}
	if err := godotenv.Load(); err != nil {
		logger.Errorf("error loading env variables: %s", err.Error())
		return
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logger.Errorf("failed to initialize db: %s", err.Error())
		return
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(moneytracker.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logger.Errorf("error occured while running http server: %s", err.Error())
		return
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
