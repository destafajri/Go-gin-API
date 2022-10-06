package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)

	//path dengan variable id
	router.GET("/books/:id", booksHandler)

	//path dengan variable id dan title
	router.GET("/books/:id/:title", booksHandlers)

	//membuat request query untuk id
	router.GET("/query", queryHandler)

	//membuat multi request query title dan price
	router.GET("/queries", queryHandlers)

	router.Run()
}

//function handler
func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name" : "Desta",
		"status" : "Belajar Golang",
	})
}

//function handler membuat path untuk id
func booksHandler(c *gin.Context){
	//mengambil parameter id
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func booksHandlers(c *gin.Context){
	//mengambil parameter id dan title
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

//function handler query untuk id
func queryHandler(c *gin.Context){
	//mengambil query id
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{"id": id})
}

//function handler query untuk title dan price
func queryHandlers(c *gin.Context){
	//mengambil query title dan price
	title := c.Query("title")
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}