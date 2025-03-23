package main

import (
	"V-Manager/internal/handlers"
	"V-Manager/internal/middleware"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	// set a middleware for url groups
	protected := r.Group("/v1")
	protected.Use(middleware.AuthMiddleware())

	// sample route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Bienvenue sur le backend de V-Manager !",
		})
	})

	r.POST("/register", handlers.Register)

	protected.GET("/metrics", handlers.GetMetrics)

	// Initializing database
	initDatabase()
	// Starting webserver
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Erreur lors du démarrage du serveur: ", err)
	}
}
