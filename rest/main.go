package main

import (
	"github.com/gin-gonic/gin"
	"temmix.com/rest-api/db"
	"temmix.com/rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run(":7070") // localhost:7070
}
