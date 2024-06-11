package routes

import (
	"github.com/gin-gonic/gin"
	"temmix.com/rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticatedRoute := server.Group("/")
	authenticatedRoute.Use(middlewares.Authenticate)
	authenticatedRoute.POST("/events", createEvent)
	authenticatedRoute.PUT("/events/:id", updateEvent)
	authenticatedRoute.DELETE("/events/:id", deleteEvent)
	authenticatedRoute.POST("/events/:id/register", registerEvent)
	authenticatedRoute.DELETE("/events/:id/register", cancelEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
