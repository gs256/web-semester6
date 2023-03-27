package orders

import (
	"lab5/products"
	"lab5/users"
)

type OrderService struct {
	repo *Repository
}

func NewOrderService(repo *Repository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (service *OrderService) CreateOrder(userId string, productIds []string) (string, error) {
	products_ := make([]products.Product, len(productIds))
	for i, pid := range productIds {
		products_[i] = products.Product{Id: pid}
	}

	order := Order{
		Id: "",
		User: users.User{
			Id: userId,
		},
		Products: products_,
	}

	id, err := service.repo.Create(&order)
	return id, err
}
