package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	router.Run(":1000")
}

//function handler
func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name" : "Desta",
		"status" : "Mencari Kerja",
	})
}