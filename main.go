package main

import (
	"apirest-go-gin/database"
	"apirest-go-gin/routes"
)

func main() {
	database.ConnecDatabase()
	routes.HandleRequests()
}
