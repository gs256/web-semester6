package orders

import (
	"lab7/products"
	"lab7/users"
	"time"
)

type Order struct {
	Id        string
	User      users.User
	Products  []products.Product
	CreatedAt time.Time
}
