package main

import (
	"V-Manager/internal/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	// sample route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Bienvenue sur le backend de V-Manager !",
		})
	})

	r.GET("/metrics", handlers.GetMetrics)

	initDatabase()

	// Starting webserver
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Erreur lors du démarrage du serveur: ", err)
	}
}
