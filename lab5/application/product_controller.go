package application

import (
	"lab5/product_repository"
	"lab5/products"
	"lab5/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	repo *product_repository.Repository
}

func (controller *ProductController) Initialize(engine *gin.Engine, repo *product_repository.Repository) {
	controller.repo = repo
	engine.GET("/products", controller.productsRoute)
	engine.GET("/products/:productId", controller.productDetailsRoute)
}

func (controller *ProductController) productsRoute(c *gin.Context) {
	p, err := controller.repo.GetAll()
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
	product, err := controller.repo.GetById(productId)

	if err != nil {
		return404(c)
	} else {
		c.HTML(http.StatusOK, "details.tmpl", product)
	}
}

func return404(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}
