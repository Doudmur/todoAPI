package main

import (
	"context"
	"todoAPI/etc/logger"
	"todoAPI/internal/database"
	"todoAPI/internal/service"
)

func main() {
	logger.SetErrorLevel()
	ctx := context.Background()
	storage, err := database.New()
	if err != nil {
		logger.Errorf(ctx, "Something wrong with connection to storage!", err)
	}

	s, err := service.New(storage)
	if err != nil {
		logger.Errorf(ctx, "Storage does not created!", err)
	}
	logger.Infof(ctx, "Server is running")
	s.Run()
}
