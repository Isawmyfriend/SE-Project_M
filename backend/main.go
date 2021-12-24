package main

import (
	"backend/config"
	"backend/router"
	"fmt"
)


func main() {

	// Config
	config.Init()
	conf := config.GetConfig()

	app := router.Router()
	
	// Start server
	app.Listen(fmt.Sprintf(":" + conf.GetString("server.port")))

}

