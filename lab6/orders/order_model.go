package orders

import (
	"lab6/products"
	"lab6/users"
	"time"
)

type OrderModel struct {
	Id string `gorm:"primaryKey;joinForeignKey:OrderId"`

	UserId string
	User   users.UserModel `gorm:"foreignKey:UserId;references:Id;constraint:OnDelete:CASCADE;"`

	Products []products.ProductModel `gorm:"many2many:order_products;constraint:OnDelete:CASCADE;"`

	CreatedAt time.Time
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
