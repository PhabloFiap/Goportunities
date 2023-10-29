package main

import (
	"github.com/PhabloFiap/Goportunities.git/config"
	"github.com/PhabloFiap/Goportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	//initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		panic(err)
	}

	//initialize router
	router.Initialize()

}
