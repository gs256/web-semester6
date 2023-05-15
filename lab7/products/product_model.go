package products

type ProductModel struct {
	Id          string `gorm:"primaryKey"`
	Name        string
	Description string
	Price       int
	ImageUrl    string
}

func (ProductModel) TableName() string {
	return "product"
}
