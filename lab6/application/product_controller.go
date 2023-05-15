package application

import (
	"lab6/products"
	"lab6/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service *products.ProductService
}

func (controller *ProductController) Initialize(engine *gin.Engine, service *products.ProductService) {
	controller.service = service
	engine.GET("/products", controller.indexRoute)
	engine.GET("/products/:productId", controller.productDetailsRoute)
	engine.GET("/cart", controller.indexRoute)
}

func (controller *ProductController) indexRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func (controller *ProductController) productsRoute(c *gin.Context) {
	p, err := controller.service.GetAllProducts()
	productsOnPage := p[:]

	page, err := strconv.Atoi(c.Query("page"))

	if err != nil || page <= 0 {
		page = 1
	}

	max, err := strconv.Atoi(c.Query("max"))

	if err != nil {
		max = len(productsOnPage)
	}

	maxPageNumber := products.MaxPageNumber(&productsOnPage, max)
	page = utils.Clamp(page, 1, maxPageNumber)

	productsOnPage = *products.ProductSlice(&productsOnPage, page, max)
	c.HTML(http.StatusOK, "products.tmpl", gin.H{"products": productsOnPage, "page": page, "pageCount": maxPageNumber})
}

func (controller *ProductController) productDetailsRoute(c *gin.Context) {
	productId := c.Param("productId")
	product, err := controller.service.GetProductById(productId)

	if err != nil {
		return404(c)
	} else {
		c.HTML(http.StatusOK, "details.tmpl", product)
	}
}

func (controller *ProductController) cartRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "cart.tmpl", nil)
}

func return404(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}
