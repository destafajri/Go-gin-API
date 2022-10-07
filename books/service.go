package books

//kontrak
type Service interface{
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(ID int, book BookRequest) (Book, error)
}

type service struct{
	respository Respository
}

//implementasi service
func NewService(respository Respository) *service{
	return &service{respository}
}

func(s*service) FindAll() ([]Book, error){
	books, err := s.respository.FindAll()
	return books, err
}

func(s*service) FindByID(ID int) (Book, error){
	book, err := s.respository.FindByID(ID)
	return book, err
}

func(s*service) Create(bookRequest BookRequest) (Book, error){
	price, _ 	:= bookRequest.Price.Int64()
	rating, _ 	:= bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()
	
	book := Book{
		Title 		: bookRequest.Title,
		Price 		: int(price),
		Description	: bookRequest.Description,
		Rating		:	int(rating),
		Discount	: int(discount),
	}

	newBook, err := s.respository.Create(book)
	return newBook, err
}

func(s*service) Update(ID int, bookRequest BookRequest) (Book, error){
	//get book
	book, err := s.respository.FindByID(ID)
	
	price, _ 	:= bookRequest.Price.Int64()
	rating, _ 	:= bookRequest.Rating.Int64()
	discount, _ := bookRequest.Discount.Int64()
	//set book
	book.Title 		= bookRequest.Title
	book.Price 		= int(price)
	book.Description= bookRequest.Description
	book.Rating		= int(rating)
	book.Discount	= int(discount)

	newBook, err := s.respository.Update(book)
	return newBook, err
}