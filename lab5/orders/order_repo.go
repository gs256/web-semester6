package orders

import (
	"fmt"
	"lab5/products"

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

func (r *Repository) GetById(id string) (*Order, error) {
	model := &OrderModel{}

	err := r.db.First(model, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	user := ToOrder(model)
	return &user, nil
}

func (r *Repository) GetAll() ([]Order, error) {
	var models []OrderModel
	err := r.db.Find(&models).Error

	if err != nil {
		return nil, err
	}

	orders := make([]Order, len(models))

	for i := 0; i < len(models); i++ {
		orders[len(models)-i-1] = ToOrder(&models[i])
	}

	return orders, nil
}

func (r *Repository) Create(order *Order) (string, error) {
	if len(order.Id) == 0 {
		order.Id = uuid.New().String()
	}
	model := ToModel(order)
	return order.Id, r.db.Create(&model).Error
}

func (r *Repository) Update(order *Order) error {
	model := ToModel(order)
	return r.db.Model(&OrderModel{}).Where("id = ?", model.Id).Updates(model).Error
}

func (r *Repository) Delete(id string) error {
	return r.db.Delete(&OrderModel{}, "id = ?", id).Error
}

func (r *Repository) Clear() error {
	return r.db.Exec(fmt.Sprintf("TRUNCATE %s;", OrderModel{}.TableName())).Error
}

func ToOrder(model *OrderModel) Order {
	productIds := make([]string, len(model.Products))
	for i, product := range model.Products {
		productIds[i] = product.Id
	}

	order := Order{
		Id:         model.Id,
		UserId:     model.User.Id,
		ProductIds: productIds,
	}

	return order
}

func ToModel(order *Order) OrderModel {
	productModels := make([]products.ProductModel, len(order.ProductIds))
	for i, productId := range order.ProductIds {
		productModels[i] = products.ProductModel{Id: productId}
	}

	model := OrderModel{
		Id:       order.Id,
		UserId:   order.UserId,
		Products: productModels,
	}

	return model
}
