package orders

import (
	"lab5/products"
	"lab5/users"
)

type Order struct {
	Id       string
	User     users.User
	Products []products.Product
}
