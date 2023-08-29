package main

import (
	"task-5-vix-btpns/database"
	"task-5-vix-btpns/router"
)

func main() {
	database.ConnectDatabase()

	router.Routers()
}