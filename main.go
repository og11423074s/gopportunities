package main

import (
	"github.com/og11423074s/gopportunities/config"
	"github.com/og11423074s/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	// init configs
	err := config.Init()

	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	// init router
	router.Initialize()
}
