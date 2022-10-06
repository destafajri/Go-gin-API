package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	router.Run(":8000")
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

/*POST*/

//membuat struct untuk menangkap data post request
type BookInput struct{
	//mengharuskan data json untuk diisi
	Title string `json:"title" binding:"required"`
	Price int	`json:"price" binding:"required,number"`
	Email string `json:"email" binding:"required,email"`

}
//function handler query untuk post
func postBookHandler(c *gin.Context){
	//membuat variable input
	var bookInput BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil{
		//validation error
		for _, e:= range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error on %s, where condition %s", e.Field(), e.ActualTag())
			c.JSON(http.StatusBadRequest, errorMessage)
		}

	}else {
		//status 201 untuk post
		c.JSON(http.StatusCreated, gin.H{
			"title": bookInput.Title,
			"price": bookInput.Price,
			"email" : bookInput.Email,
		})
	}
}