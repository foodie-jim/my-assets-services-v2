package main

import (
	"log"
	"net/http"
	"os"
	
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	} else {
		log.Println("$PORT is " + port)
	}

	r := gin.Default()
	r.Use(gin.Logger())
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":" + port)
}