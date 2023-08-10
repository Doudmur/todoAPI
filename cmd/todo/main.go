package main

import (
	"context"
	"os"
	"syscall"
	"todoAPI/closer"
	"todoAPI/etc/logger"
	"todoAPI/internal/database"
	"todoAPI/internal/service"
)

func main() {
	logger.SetErrorLevel(4)
	ctx := context.Background()
	storage, err := database.New()
	logger.Infof(ctx, "Server")
	if err != nil {
		logger.Errorf(ctx, "Something wrong with connection to storage!", err)
	}

	s, err := service.New(storage)
	if err != nil {
		logger.Errorf(ctx, "Storage does not created!", err)
	}
	logger.Infof(ctx, "Server is running")
	s.Run()

	cl := closer.New(syscall.SIGINT, syscall.SIGTERM)
	cl.Add(s.Shutdown)

	cl.Close()
	os.Exit(0)
}
