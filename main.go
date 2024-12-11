package main

import (
	"github.com/jeanpatrickm/job-manager/config"
	"github.com/jeanpatrickm/job-manager/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// initialize config	
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}
	
	router.InitializeRouter()
}