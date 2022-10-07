package books

import "gorm.io/gorm"

//kontrak
type Respository interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book Book) (Book, error)
	Update(book Book) (Book, error)
	Delete(book Book) (Book, error)
}

type respository struct {
	db *gorm.DB

}

//implementasi repository
func NewRepository(db *gorm.DB) *respository{
	return &respository{db}
}

func(r*respository) FindAll() ([]Book, error){
	var books []Book
	err := r.db.Find(&books).Error
	return books, err
}

func(r*respository) FindByID(ID int) (Book, error){
	var book Book
	err := r.db.Find(&book, ID).Error
	return book, err
}

func(r*respository) Create(book Book) (Book, error){
	err := r.db.Create(&book).Error
	return book, err
}

func(r*respository) Update(book Book) (Book, error){
	err := r.db.Save(&book).Error
	return book, err
}

func(r*respository) Delete(book Book) (Book, error){
	err := r.db.Delete(&book).Error
	return book, err
}