package application

import (
	"lab5/database"
	"lab5/products"
	"lab5/users"

	"github.com/gin-gonic/gin"
)

type Application struct {
	engine         *gin.Engine
	productRepo    *products.Repository
	userRepo       *users.Repository
	productService *products.ProductService
	userService    *users.UserService
}

func New() *Application {
	return &Application{}
}

func (app *Application) Run(engine *gin.Engine) {
	app.engine = engine
	app.productRepo, _ = products.NewRepository(database.GetDsn())
	app.userRepo, _ = users.NewRepository(database.GetDsn())
	app.productService = products.NewProductService(app.productRepo)
	app.userService = users.NewUserService(app.userRepo)
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
	apiController.Initialize(app.engine, app.productService, app.userService)
}
