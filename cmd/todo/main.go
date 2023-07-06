package main

import (
	"todoAPI/etc/logger"
	"todoAPI/internal/database"
	"todoAPI/internal/service"
)

func main() {
	l := logger.NewLogger()
	storage, err := database.New()
	if err != nil {
		l.Error("Something wrong with connection to storage!", err)
	}

	s, err := service.New(storage)
	if err != nil {
		l.Error("Storage does not created!", err)
	}
	l.Info("Server is running")
	s.Run()
}
