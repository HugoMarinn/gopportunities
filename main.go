package main

import (
	"github.com/HugoMarinn/gopportunities/config"
	"github.com/HugoMarinn/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("config inicialization error: %v", err)
		return 
	}

	router.Initialize()
}
