package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"pustaka-api/handler"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//database connection
	dsn := "root:@tcp(127.0.0.1:3306)/intern_privy?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Database connection error")
	}
	fmt.Println("Database Connected....................................................")

	//router default setting
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
	router.Run()


}