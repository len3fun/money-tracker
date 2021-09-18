package main

import (
	"github.com/joho/godotenv"
	"github.com/len3fun/money-tracker/internal/handler"
	"github.com/len3fun/money-tracker/internal/repository"
	"github.com/len3fun/money-tracker/internal/server"
	"github.com/len3fun/money-tracker/internal/service"
	"github.com/len3fun/money-tracker/pkg/logger"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		logger.Errorf("error initializing config: %s", err.Error())
		return
	}
	if err := godotenv.Load(); err != nil {
		logger.Errorf("error loading env variables: %s", err.Error())
		return
	}

	logger.InitLogger()

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

	go func() {
		srv := new(server.Server)
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logger.Errorf("error occurred while running http server: %s", err.Error())
			return
		}
	}()

	logger.Info("Server has started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("Server is stopping")
	if err := db.Close(); err != nil {
		logger.Errorf("close db connection error: %s", err.Error())
	}

	logger.Info("Server is stopped")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
