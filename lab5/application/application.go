package application

import (
	"lab5/database"
	"lab5/products"

	"github.com/gin-gonic/gin"
)

type Application struct {
	productRepo    *products.Repository
	engine         *gin.Engine
	productService *products.ProductService
}

func New() *Application {
	return &Application{}
}

func (app *Application) Run(engine *gin.Engine) {
	app.engine = engine
	app.productRepo, _ = products.NewRepository(database.GetDsn())
	app.productService = products.NewProductService(app.productRepo)
	app.initializeControllers()
}

func (app *Application) initializeControllers() {
	testController := TestController{}
	productController := ProductController{}
	adminController := AdminController{}
	apiController := ApiController{}

	testController.Initialize(app.engine, app.productRepo)
	productController.Initialize(app.engine, app.productRepo)
	adminController.Initialize(app.engine, app.productRepo)
	apiController.Initialize(app.engine, app.productService)
}
