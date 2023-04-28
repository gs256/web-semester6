package application

import (
	"lab6/database"
	"lab6/orders"
	"lab6/products"
	"lab6/users"

	"github.com/gin-gonic/gin"
)

type Application struct {
	engine         *gin.Engine
	productRepo    *products.Repository
	userRepo       *users.Repository
	orderRepo      *orders.Repository
	productService *products.ProductService
	userService    *users.UserService
	orderService   *orders.OrderService
}

func New() *Application {
	return &Application{}
}

func (app *Application) Run(engine *gin.Engine) {
	app.engine = engine
	app.productRepo, _ = products.NewRepository(database.GetDsn())
	app.userRepo, _ = users.NewRepository(database.GetDsn())
	app.orderRepo, _ = orders.NewRepository(database.GetDsn())
	app.productService = products.NewProductService(app.productRepo)
	app.userService = users.NewUserService(app.userRepo)
	app.orderService = orders.NewOrderService(app.orderRepo, app.userRepo, app.productRepo)
	app.initializeControllers()
}

func (app *Application) initializeControllers() {
	testController := TestController{}
	productController := ProductController{}
	adminController := AdminController{}
	apiController := ApiController{}

	testController.Initialize(app.engine, app.productRepo)
	productController.Initialize(app.engine, app.productService)
	adminController.Initialize(app.engine, app.productRepo, app.orderService, app.userService)
	apiController.Initialize(app.engine, app.productService, app.userService, app.orderService)
}
