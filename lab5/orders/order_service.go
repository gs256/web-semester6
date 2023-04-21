package orders

import (
	"fmt"
	"lab5/products"
	"lab5/users"
)

type OrderService struct {
	orderRepo   *Repository
	userRepo    *users.Repository
	productRepo *products.Repository
}

func NewOrderService(orderRepo *Repository, userRepo *users.Repository, productRepo *products.Repository) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		userRepo:    userRepo,
		productRepo: productRepo,
	}
}

func (service *OrderService) CreateOrder(userId string, productIds []string) (string, error) {
	_, err := service.userRepo.GetById(userId)

	if err != nil {
		return "", fmt.Errorf("No user with id %s", userId)
	}

	for _, productId := range productIds {
		_, err = service.productRepo.GetById(productId)
		if err != nil {
			return "", fmt.Errorf("No product with id %s", productId)
		}
	}

	products_ := make([]products.Product, len(productIds))
	for i := 0; i < len(products_); i++ {
		products_[i] = products.Product{Id: productIds[i]}
	}

	order := Order{
		Id:       "",
		User:     users.User{Id: userId},
		Products: products_,
	}

	id, err := service.orderRepo.Create(&order)
	return id, err
}

func (service *OrderService) GetAllOrders() ([]Order, error) {
	orders, err := service.orderRepo.GetAll()
	return orders, err
}

func (service *OrderService) Clear() error {
	return service.orderRepo.Clear()
}
