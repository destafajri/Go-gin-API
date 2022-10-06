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
/****************Database***************/
	//database connection
	dsn := "root:@tcp(127.0.0.1:3306)/intern_privy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Database connection error")
	}

	//database auto migrate
	db.AutoMigrate(&books.Book{})

	/*****CRUD database***/
	/*Create*/
	book := books.Book{}
	book.Title = "Ekoju"
	book.Price = 50000
	book.Discount = 100
	book.Rating = 10
	book.Description = "Tips cuan dari Eko alias Lord Oura nih"
	er := db.Debug().Create(&book).Error
	if er != nil{
		log.Fatal("Create request error")
	}

	/*Read*/
	//read first book
	var book1 books.Book
	er1 := db.Debug().First(&book1).Error
	if er1 != nil{
		log.Fatal("read request error")
	}
	fmt.Println("Books Title ", book1.Title)

	//read book with primary key
	var book2 books.Book
	er2 := db.Debug().First(&book2, 3).Error
	if er2 != nil{
		log.Fatal("read request error")
	}
	fmt.Println("Books Title ", book2.Title)

	//read last book
	var book3 books.Book
	er3 := db.Debug().Last(&book3).Error
	if er3 != nil{
		log.Fatal("read request error")
	}
	fmt.Println("Books Title ", book3.Title)
	fmt.Printf("Books object ===>%v ", book3)

	//read all object
	// var books []books.Book
	// errs := db.Debug().Find(&books).Error
	// if errs != nil{
	// 	log.Fatal("read request error")
	// }
	// for _, p := range books{
	// 	fmt.Println("Books Title ", p.Title)
	// 	fmt.Printf("Books object ===> %v ", p)
	// }

	//read book with condition
	// var books2 []books.Book
	// errs2 := db.Debug().Where("price > 500000").Find(&books).Error
	// if errs2 != nil{
	// 	log.Fatal("read request error")
	// }
	// for _, p := range books{
	// 	fmt.Println("Books Title ", p.Title)
	// 	fmt.Printf("Books object ===> %v ", p)
	// }
	
	/*Update*/
	//select the data
	var uBook books.Book
	errU := db.Debug().First(&uBook, 5).Error
	if errU != nil{
		log.Fatal("read request error")
	}
	fmt.Println("Books Title ", uBook.Title)
	fmt.Printf("Books object ===> %v ", uBook)
	//and then update
	uBook.Title = "Eko Up to date lagi nih"
	db.Debug().Save(&uBook)
	fmt.Println("Books Update ", uBook)

	/*Delete*/
	//Select the data
	var dbook books.Book
	errd := db.Debug().Last(&dbook).Error
	if errd != nil{
		log.Fatal("read request error")
	}
	//And the delete
	db.Debug().Delete(&dbook)


/*********API Request**********/
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