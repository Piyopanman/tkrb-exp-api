package main

import (
	"tkrb-exp-api/src/logging"
	"tkrb-exp-api/src/server"
)

func main() {
	logging.InitLogger(logging.Dev)
	// db.Init()
	server.Init()
}
