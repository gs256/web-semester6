package database

type ProductModel struct {
	Pk          uint `gorm:"primarykey"`
	Id          string
	Name        string
	Description string
	Price       int
	ImageUrl    string
}

func (ProductModel) TableName() string {
	return "product"
}
