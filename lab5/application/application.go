package application

import (
	"lab5/database"
	"lab5/products"

	"github.com/gin-gonic/gin"
)

type Application struct {
	repo   *products.Repository
	engine *gin.Engine
}

func New() *Application {
	return &Application{}
}

func (app *Application) Run(engine *gin.Engine) {
	app.engine = engine
	app.repo, _ = products.NewRepository(database.GetDsn())
	app.initializeControllers()
}

func (app *Application) initializeControllers() {
	testController := TestController{}
	productController := ProductController{}
	adminController := AdminController{}

	testController.Initialize(app.engine, app.repo)
	productController.Initialize(app.engine, app.repo)
	adminController.Initialize(app.engine, app.repo)
}
