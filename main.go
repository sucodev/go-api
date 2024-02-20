package main

import (
	"github.com/sucodev/go-api/config"
	"github.com/sucodev/go-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Intialize Router
	router.Initialize()
}
