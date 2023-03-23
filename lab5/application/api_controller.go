package application

import (
	"lab5/products"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
	productService *products.ProductService
}

func (controller *ApiController) Initialize(engine *gin.Engine, productService *products.ProductService) {
	controller.productService = productService
	engine.GET("/api/products", controller.productsRoute)
}

func (controller *ApiController) productsRoute(c *gin.Context) {
	products, err := controller.productService.GetAllProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get products"})
	}

	c.JSON(http.StatusOK, products)
}
