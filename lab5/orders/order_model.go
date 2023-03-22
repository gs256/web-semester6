package orders

import (
	"lab5/products"
	"lab5/users"
)

type OrderModel struct {
	Pk       uint `gorm:"primarykey"`
	Id       string
	UserId   uint
	User     users.UserModel         `gorm:"foreignKey:UserId"`
	Products []products.ProductModel `gorm:"many2many:product_orders;"`
}

func (OrderModel) TableName() string {
	return "order"
}
