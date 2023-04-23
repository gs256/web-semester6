package orders

import (
	"lab5/products"
	"lab5/users"
	"time"
)

type Order struct {
	Id        string
	User      users.User
	Products  []products.Product
	CreatedAt time.Time
}
