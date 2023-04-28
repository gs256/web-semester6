package orders

import (
	"fmt"
	"lab6/products"
	"lab6/users"
	"time"
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
		Id:        "",
		User:      users.User{Id: userId},
		Products:  products_,
		CreatedAt: time.Now(),
	}

	id, err := service.orderRepo.Create(&order)
	return id, err
}

func (service *OrderService) GetAllOrders() ([]Order, error) {
	orders, err := service.orderRepo.GetAll()
	return orders, err
}

func (service *OrderService) GetWithinTimespan(start time.Time, end time.Time) ([]Order, error) {
	orders, err := service.orderRepo.GetWithinTimespan(start, end)
	return orders, err
}

func (service *OrderService) Clear() error {
	return service.orderRepo.Clear()
}

func (service *OrderService) GetUserOrderHistory(userId string) ([]products.Product, error) {
	_, err := service.userRepo.GetById(userId)

	if err != nil {
		return nil, fmt.Errorf("No user with id %s", userId)
	}

	orders, err := service.orderRepo.GetByUserId(userId)

	if err != nil {
		return nil, fmt.Errorf("No user with id %s", userId)
	}

	products_ := make([]products.Product, 0)

	for _, order := range orders {
		for _, product := range order.Products {
			products_ = append(products_, product)
		}
	}

	return products_, nil
}
