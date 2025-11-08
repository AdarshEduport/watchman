package types

import "time"


type IUserStore interface {
	GetUserByEmail(email string)(*User ,error)
	CreateUser(user User)error
	GetUserById(id int)(*User,error)
}
type IProductStore interface {
	GetProducts() ([]Product, error)
}


type AuthUserPayload struct{
	Name string `json:"name" validate:"required"` 
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	
}

type User struct{
	Id int `json:"Id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type Product struct{
	Id int `json:"Id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
	Quantity int `json:"quantity"`
	Image string `json:"image"`

}