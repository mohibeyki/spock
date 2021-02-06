package main

import (
	"fmt"
	"log"

	"github.com/mohibeyki/spock/pkg/config"
	"github.com/mohibeyki/spock/pkg/database"
	"github.com/mohibeyki/spock/pkg/router"
)

func init() {
	config.Init()
	database.Init()
}

func main() {
	config := config.GetConfig()
	rootRouter := router.Init(database.GetDB())

	socketString := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	log.Printf("Server is running at %s", socketString)
	rootRouter.Run(socketString)
}
