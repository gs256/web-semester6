package application

import (
	"lab5/products"
	"lab5/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
	productService *products.ProductService
	userService    *users.UserService
}

func (controller *ApiController) Initialize(engine *gin.Engine, productService *products.ProductService, userService *users.UserService) {
	controller.productService = productService
	controller.userService = userService
	engine.GET("/api/products", controller.productsRoute)
	engine.GET("/api/users", controller.usersRoute)
}

func (controller *ApiController) productsRoute(c *gin.Context) {
	products, err := controller.productService.GetAllProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get products"})
	}

	c.JSON(http.StatusOK, products)
}

func (controller *ApiController) usersRoute(c *gin.Context) {
	users, err := controller.userService.GetAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get users"})
	}

	c.JSON(http.StatusOK, users)
}
