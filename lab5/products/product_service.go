package products

import (
	"lab5/utils"
	"math"
)

type ProductService struct {
	repo *Repository
}

func NewProductService(repo *Repository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (service *ProductService) GetAllProducts() ([]Product, error) {
	if service == nil {
		panic("service is nil")
	}
	if service.repo == nil {
		panic("repo is nil")
	}
	products, err := service.repo.GetAll()
	return products, err
}

func MaxPageNumber(products *[]Product, itemsOnPage int) int {
	if len(*products) <= 0 {
		return 1
	}

	if itemsOnPage <= 0 || itemsOnPage >= len(*products) {
		return 1
	}

	numberOfPages := float64(len(*products)) / float64(itemsOnPage)
	return int(math.Ceil(numberOfPages))
}

func ProductSlice(products *[]Product, page int, max int) *[]Product {
	var start int
	var end int

	if page <= 0 {
		page = 1
	}

	if max <= 0 {
		return &[]Product{}
	}

	start = (page - 1) * max
	end = start + max

	end = utils.Clamp(end, 0, len(*products))
	start = utils.Clamp(start, 0, end)

	slice := (*products)[start:end]
	return &slice
}
