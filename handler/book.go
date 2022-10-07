package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"pustaka-api/books"
)

//mengakses service
type bookHandler struct{
	bookService books.Service
}

func NewBookHandler(bookService books.Service) *bookHandler{
	return &bookHandler{bookService}
}

/*GET*/
//function handler
func (handler *bookHandler)RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name" : "Desta",
		"status" : "Belajar Golang",
	})
}

//function handler membuat path untuk id
func (handler *bookHandler)BooksHandler(c *gin.Context){
	//mengambil parameter id
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (handler *bookHandler)BooksHandlers(c *gin.Context){
	//mengambil parameter id dan title
	id := c.Param("id")
	title := c.Param("title")

	c.JSON(http.StatusOK, gin.H{"id": id, "title": title})
}

//function handler query untuk id
func (handler *bookHandler)QueryHandler(c *gin.Context){
	//mengambil query id
	id := c.Query("id")

	c.JSON(http.StatusOK, gin.H{"id": id})
}

//function handler query untuk title dan price
func (handler *bookHandler)QueryHandlers(c *gin.Context){
	//mengambil query title dan price
	title := c.Query("title")
	price := c.Query("price")
	
	c.JSON(http.StatusOK, gin.H{"title": title, "price": price})
}


/*POST*/
//function handler query untuk post
func (handler *bookHandler)PostBookHandler(c *gin.Context){
	//membuat variable input
	var bookInput books.BookRequest

	err := c.ShouldBindJSON(&bookInput)
	if err != nil{
		//slice err
		errorMessages := []string{}
		//validation error
		for _, e:= range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error on %s, where condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
			c.JSON(http.StatusBadRequest, gin.H{
				"errors" : errorMessages,
			})
		}

	}else {
		//status 201 untuk post
		book, err := handler.bookService.Create(bookInput)
		if err !=nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errors" : err,
			})
		}else{
			c.JSON(http.StatusCreated, gin.H{
				"data" : book,
			})
		}
	}
}