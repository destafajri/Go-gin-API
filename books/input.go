package books

import (
	"encoding/json"

)

/*POST API*/
//membuat struct untuk menangkap data post request
type BookInput struct{
	//mengharuskan data json untuk diisi
	Title string `json:"title" binding:"required"`
	Price json.Number	`json:"price" binding:"required,number"`
	Email string `json:"email" binding:"required,email"`

}
