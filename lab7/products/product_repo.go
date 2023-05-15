package products

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(dsn string) (*Repository, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) GetById(id string) (*Product, error) {
	productModel := &ProductModel{}

	err := r.db.First(productModel, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	product := ToProduct(productModel)
	return &product, nil
}

func (r *Repository) GetAll() ([]Product, error) {
	var productModels []ProductModel
	err := r.db.Find(&productModels).Error

	if err != nil {
		return nil, err
	}

	products := make([]Product, len(productModels))

	for i := 0; i < len(productModels); i++ {
		products[len(productModels)-i-1] = ToProduct(&productModels[i])
	}

	return products, err
}

func (r *Repository) Create(product *Product) error {
	if len(product.Id) == 0 {
		product.Id = uuid.New().String()
	}
	model := ToModel(product)
	return r.db.Create(&model).Error
}

func (r *Repository) Update(product *Product) error {
	model := ToModel(product)
	return r.db.Model(&ProductModel{}).Where("id = ?", model.Id).Updates(model).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.Delete(&ProductModel{}, "id = ?", id).Error
}

func (r *Repository) Clear() error {
	return r.db.Exec(fmt.Sprintf("TRUNCATE %s;", ProductModel{}.TableName())).Error
}

func ToProduct(model *ProductModel) Product {
	product := Product{
		Id:          model.Id,
		Name:        model.Name,
		Description: model.Description,
		Price:       model.Price,
		ImageUrl:    model.ImageUrl,
	}

	return product
}

func ToModel(product *Product) ProductModel {
	model := ProductModel{
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		ImageUrl:    product.ImageUrl,
	}

	return model
}
