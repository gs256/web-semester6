package application

import (
	"lab5/database"
	"lab5/product_repository"
	"lab5/products"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TestController struct {
	repo *product_repository.Repository
}

func (controller *TestController) Initialize(engine *gin.Engine, repo *product_repository.Repository) {
	controller.repo = repo
	engine.GET("/test", controller.testRoute)
	engine.GET("/test/create-random", controller.createRandomRoute)
	engine.GET("/test/fill-with-test-data", controller.fillDbWithTestProducts)
	engine.GET("/test/clear", controller.clearRoute)
}

func (controller *TestController) testRoute(c *gin.Context) {
	products, err := controller.repo.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (controller *TestController) createRandomRoute(c *gin.Context) {
	product := products.Product{
		Id:          uuid.New().String(),
		Name:        randomString(6),
		Description: randomString(50),
		Price:       rand.Intn(5000) + 1000,
		ImageUrl:    "https://dummyimage.com/512",
	}

	err := controller.repo.Create(&product)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't create a new product"})
	} else {
		c.Status(http.StatusOK)
	}
}

func (controller *TestController) clearRoute(c *gin.Context) {
	err := controller.repo.Clear()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Table truncate failed"})
	} else {
		c.Status(http.StatusOK)
	}
}

func (controller *TestController) fillDbWithTestProducts(c *gin.Context) {
	repo, _ := product_repository.New(database.GetDsn())

	for _, product := range ProductList {
		repo.Create(&product)
	}
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
