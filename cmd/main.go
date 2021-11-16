package main

import (
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
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

	const migrationsVersion = 1

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		logger.Errorf("failed init driver for migration tool: %s", err.Error())
		return
	}

	m, err := migrate.NewWithDatabaseInstance("file://./schema", "postgres", driver)
	if err != nil {
		logger.Errorf("failed init migrations tool: %s", err.Error())
		return
	}

	if err := applyMigrations(migrationsVersion, m); err != nil {
		logger.Errorf("failed to apply migrations: %s", err.Error())
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

func applyMigrations(version uint, m *migrate.Migrate) error {
	currentVersion, _, err := m.Version()
	if err != nil {
		switch err {
		case migrate.ErrNilVersion:
			logger.Info("DB doesn't have any migrations, try to apply latest version")
			if err := m.Migrate(version); err != nil {
				return err
			}
			currentVersion = version
		case err:
			return err
		}
	}

	logger.Debugf("Current migrations version: %d", currentVersion)
	if currentVersion != version {
		logger.Info("DB has outdated migrations version, try to apply latest version")
		if err := m.Migrate(version); err != nil {
			return err
		}
	}

	return nil
}
