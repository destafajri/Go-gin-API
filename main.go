package main

import (
	"log"
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

	//Post request
	router.POST("/books", postBookHandler)

	router.Run()
}

/*GET*/

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
	//mengembalikan nilai json
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func booksHandlers(c *gin.Context){
	//mengambil parameter id dan title
	id := c.Param("id")
	title := c.Param("title")
	//mengembalikan nilai json
	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

//function handler query untuk id
func queryHandler(c *gin.Context){
	//mengambil query id
	id := c.Query("id")
	//mengembalikan nilai json
	c.JSON(http.StatusOK, gin.H{"id": id})
}

//function handler query untuk title dan price
func queryHandlers(c *gin.Context){
	//mengambil query title dan price
	title := c.Query("title")
	price := c.Query("price")
	
	//mengembalikan nilai json
	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}

/*POST*/

//membuat struct untuk menangkap data post request
type BookInput struct{
	Title string
	Price int
	SubTitle string `json:"sub_title"`
}
//function handler query untuk post
func postBookHandler(c *gin.Context){
	//membuat variable input
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil{
		log.Fatal(err)
	}
	//mengembalikan nilai json
	//status 201 untuk post
	c.JSON(http.StatusCreated, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
		"sub_title": bookInput.SubTitle,
	})
}