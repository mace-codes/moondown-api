package main

import (
	"github.com/mace-codes/moondown-api/internal/app"
	"go.uber.org/zap"
)

func main() {
	if err := app.Run(); err != nil {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		logger.Fatal("Failed to start moondown-api", zap.Error(err))
	}
}
