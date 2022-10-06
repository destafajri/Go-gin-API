package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name" : "Desta",
			"bio" : "Programmer",
		})
	})

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name" : "Desta",
			"status" : "Mencari Kerja",
		})
	})

	router.Run()
}