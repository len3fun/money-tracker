package main

import (
	moneytracker "github.com/len3fun/money-tracker"
	"github.com/len3fun/money-tracker/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(moneytracker.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}


