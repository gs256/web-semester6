package products

import (
	"errors"
	"lab3/utils"
	"math"
)

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

func GetProductById(id string) (*Product, error) {
	for _, product := range ProductList {
		if product.Id == id {
			return &product, nil
		}
	}

	return nil, errors.New("Couldn't find product with this id")
}
