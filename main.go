package main

import (
	"github.com/gin-gonic/gin"
	
	"pustaka-api/handler"
)

func main() {
	router := gin.Default()

	//versioning v1
	v1 := router.Group("/v1")

	router.GET("/", handler.RootHandler)
	//v1 path for root request
	v1.GET("/", handler.RootHandler)
	//path dengan variable id
	router.GET("/books/:id", handler.BooksHandler)
	//path dengan variable id dan title
	router.GET("/books/:id/:title", handler.BooksHandlers)
	//membuat request query untuk id
	router.GET("/query", handler.QueryHandler)
	//membuat multi request query title dan price
	router.GET("/queries", handler.QueryHandlers)

	//Post request
	router.POST("/books", handler.PostBookHandler)

	//server
	router.Run(":8000")
}