package orders

import (
	"lab6/products"
	"lab6/users"
	"time"
)

type Order struct {
	Id        string
	User      users.User
	Products  []products.Product
	CreatedAt time.Time
}
