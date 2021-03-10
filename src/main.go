package main

import (
	"touken-exp/src/db"
	"touken-exp/src/server"
)

func main() {
	db.Init()
	server.Init()
}
