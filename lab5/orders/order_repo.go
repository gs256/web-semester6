package orders

import (
	"fmt"
	"lab5/products"
	"lab5/users"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	var orderModels []OrderModel
	err := r.db.Preload(clause.Associations).Find(&orderModels).Error

	if err != nil {
		return nil, err
	}

	orders := make([]Order, len(orderModels))

	for i := 0; i < len(orderModels); i++ {
		orders[len(orderModels)-i-1] = ToOrder(&orderModels[i])
	}

	return orders, nil
}

func (r *Repository) GetWithinTimespan(start time.Time, end time.Time) ([]Order, error) {
	var orderModels []OrderModel
	err := r.db.Where("created_at BETWEEN ? AND ?", start, end).Preload(clause.Associations).Find(&orderModels).Error

	if err != nil {
		return nil, err
	}

	orders := make([]Order, len(orderModels))

	for i := 0; i < len(orderModels); i++ {
		orders[len(orderModels)-i-1] = ToOrder(&orderModels[i])
	}

	return orders, nil
}

func (r *Repository) GetByUserId(userId string) ([]Order, error) {
	var orderModels []OrderModel
	err := r.db.Where("user_id = ?", userId).Preload(clause.Associations).Find(&orderModels).Error

	if err != nil {
		return nil, err
	}

	orders := make([]Order, len(orderModels))

	for i := 0; i < len(orderModels); i++ {
		orders[len(orderModels)-i-1] = ToOrder(&orderModels[i])
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
	err := r.db.Exec(fmt.Sprintf("TRUNCATE \"%s\";", OrderProductsModel{}.TableName())).Error

	if err != nil {
		return err
	}

	return r.db.Exec(fmt.Sprintf("TRUNCATE \"%s\";", OrderModel{}.TableName())).Error
}

func ToOrder(model *OrderModel) Order {
	products_ := make([]products.Product, len(model.Products))
	for i, product := range model.Products {
		products_[i] = products.ToProduct(&product)
	}

	order := Order{
		Id:        model.Id,
		User:      users.ToUser(&model.User),
		Products:  products_,
		CreatedAt: model.CreatedAt,
	}

	return order
}

func ToModel(order *Order) OrderModel {
	productModels := make([]products.ProductModel, len(order.Products))
	for i, product := range order.Products {
		productModels[i] = products.ProductModel{Id: product.Id}
	}

	model := OrderModel{
		Id:       order.Id,
		UserId:   order.User.Id,
		Products: productModels,
	}

	return model
}
