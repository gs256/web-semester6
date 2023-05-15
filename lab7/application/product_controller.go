package application

import (
	"lab7/products"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service *products.ProductService
}

func (controller *ProductController) Initialize(engine *gin.Engine, service *products.ProductService) {
	controller.service = service
	engine.GET("/products", controller.indexRoute)
	engine.GET("/products/:productId", controller.indexRoute)
	engine.GET("/cart", controller.indexRoute)
}

func (controller *ProductController) indexRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}
