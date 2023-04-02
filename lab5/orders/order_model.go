package orders

import (
	"lab5/products"
	"lab5/users"
)

type OrderModel struct {
	Id string `gorm:"primaryKey"`

	UserId string
	User   users.UserModel `gorm:"foreignKey:UserId;associationForeignKey:Id"`

	Products []products.ProductModel `gorm:"many2many:order_products"`
}

func (OrderModel) TableName() string {
	return "order"
}
