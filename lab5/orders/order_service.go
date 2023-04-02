package orders

type OrderService struct {
	repo *Repository
}

func NewOrderService(repo *Repository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (service *OrderService) CreateOrder(userId string, productIds []string) (string, error) {
	order := Order{
		Id:         "",
		UserId:     userId,
		ProductIds: productIds,
	}

	id, err := service.repo.Create(&order)
	return id, err
}
