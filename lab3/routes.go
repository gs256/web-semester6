package main

import (
	"lab3/products"
	"lab3/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	r.GET("/home", homeRoute)
	r.GET("/products", productsRoute)
	r.GET("/products/:productId", productDetailsRoute)
}

func homeRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Main website",
	})
}

func productsRoute(c *gin.Context) {
	productsOnPage := products.ProductList[:]

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

func productDetailsRoute(c *gin.Context) {
	productId := c.Param("productId")
	product, err := products.GetProductById(productId)

	if err != nil {
		return404(c)
	} else {
		c.HTML(http.StatusOK, "details.tmpl", product)
	}
}

func return404(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}
