package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
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
	go s.Run()

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)

	<-shutdownChan
	logger.Infof(ctx, "Server shutting down")
	if err = s.Shutdown(); err != nil {
		logger.Errorf(ctx, "Close error", err)
	}
	logger.Infof(ctx, "Server gracefully shutting down")
	os.Exit(0)
}
