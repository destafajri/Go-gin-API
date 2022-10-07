package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"pustaka-api/books"
	"pustaka-api/handler"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Welcome to my API....")

/****************Database***************/
	//database connection
	dsn := "root:@tcp(127.0.0.1:3306)/intern_privy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Database connection error")
	}

	//database auto migrate
	db.AutoMigrate(&books.Book{})

	/*Layer Repository & Service
	apps logic=>main->Handler/Controler->Service->repository->db->mysql
	*/
	//Repository
	bookRepository := books.NewRepository(db)
	//Service
	bookService := books.NewService(bookRepository)
	//handler/controler
	bookHandler := handler.NewBookHandler(bookService)


/*********API Request**********/
	//router default setting
	router := gin.Default()
	//versioning v1
	v1 := router.Group("/v1")

	//router path
	router.GET("/", bookHandler.RootHandler)
	//v1 path for root request
	v1.GET("/", bookHandler.RootHandler)

	//Get all books request
	router.GET("/books", bookHandler.GetBooksHandler)
	router.GET("/books/", bookHandler.GetBooksHandler)
	//Get single request books with id
	router.GET("/books/:id", bookHandler.GetBookHandler)

	//Update request
	router.PUT("/books/:id", bookHandler.PutBookHandler)

	//Post request
	router.POST("/books", bookHandler.PostBookHandler)

	//Delete request
	router.DELETE("/books/:id", bookHandler.DeleteBookHandler)

	//server
	router.Run()


}