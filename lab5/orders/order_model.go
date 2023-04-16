package orders

import (
	"lab5/products"
	"lab5/users"
)

type OrderModel struct {
	Id string `gorm:"primaryKey;joinForeignKey:OrderId"`

	UserId string
	User   users.UserModel `gorm:"foreignKey:UserId;associationForeignKey:Id"`

	Products []products.ProductModel `gorm:"many2many:order_products"`
}

func (OrderModel) TableName() string {
	return "order"
}

type OrderProductsModel struct {
	OrderModelId   string `gorm:"primaryKey"`
	ProductModelId string `gorm:"primaryKey"`
}

func (OrderProductsModel) TableName() string {
	return "order_products"
}
