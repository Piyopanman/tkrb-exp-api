package main

import (
	"touken-exp/src/logging"
	"touken-exp/src/server"
)

func main() {
	logging.InitLogger(logging.Dev)
	// db.Init()
	server.Init()
}
