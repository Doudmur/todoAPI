package main

import (
	"log"
	"todoAPI/internal/database"
	"todoAPI/internal/service"
)

func main() {
	storage, err := database.New()
	if err != nil {
		log.Fatal(err)
	}

	s, err := service.New(storage)
	if err != nil {
		log.Fatal(err)
	}

	s.Run()
}
