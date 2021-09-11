package main

import (
	moneytracker "github.com/len3fun/money-tracker"
	"log"
)

func main() {
	srv := new(moneytracker.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}


