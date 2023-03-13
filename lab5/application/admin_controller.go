package application

import (
	"encoding/json"
	"lab5/product_repository"
	"lab5/products"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	repo *product_repository.Repository
}

func (controller *AdminController) Initialize(engine *gin.Engine, repo *product_repository.Repository) {
	controller.repo = repo
	engine.GET("/admin/products", controller.adminRoute)
	engine.GET("/admin/products/delete/:id", controller.deleteRoute)
	engine.GET("/admin/products/create", controller.createFormRoute)
	engine.POST("/admin/products/create", controller.createProductRoute)
	engine.GET("/admin/products/edit/:id", controller.editFormRoute)
	engine.POST("/admin/products/edit/:id", controller.editRoute)
}

func (controller *AdminController) adminRoute(c *gin.Context) {
	products, err := controller.repo.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get products"})
	} else {
		c.HTML(http.StatusOK, "admin/admin.tmpl", gin.H{"products": products})
	}
}

func (controller *AdminController) deleteRoute(c *gin.Context) {
	id := c.Param("id")

	err := controller.repo.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete item with such id"})
	} else {
		c.Redirect(http.StatusPermanentRedirect, "/admin/products")
	}
}

func (controller *AdminController) editFormRoute(c *gin.Context) {
	id := c.Param("id")

	product, err := controller.repo.GetById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get item with such id"})
	} else {
		c.HTML(http.StatusOK, "admin/edit.tmpl", gin.H{"product": product})
	}
}

func (controller *AdminController) createFormRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/create.tmpl", gin.H{})
}

type ProductRequestDto struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
}

func (controller *AdminController) createProductRoute(c *gin.Context) {
	var productDto ProductRequestDto

	err := json.NewDecoder(c.Request.Body).Decode(&productDto)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	product := products.Product{
		Id:          "",
		Name:        productDto.Name,
		Description: productDto.Description,
		Price:       productDto.Price,
		ImageUrl:    productDto.ImageUrl,
	}

	controller.repo.Create(&product)
}

func (controller *AdminController) editRoute(c *gin.Context) {
	id := c.Param("id")
	var productDto ProductRequestDto

	err := json.NewDecoder(c.Request.Body).Decode(&productDto)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	product := products.Product{
		Id:          id,
		Name:        productDto.Name,
		Description: productDto.Description,
		Price:       productDto.Price,
		ImageUrl:    productDto.ImageUrl,
	}

	err = controller.repo.Update(&product)

	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusBadRequest)
	}
}
