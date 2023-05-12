package main

import (
	"lab6/application"
	"lab6/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// migrate()
	engine := gin.Default()
	application.New().Run(engine)

	engine.Static("./static", "./static/")
	engine.LoadHTMLGlob("templates/**/*")
	engine.Static("/images/products", "./resources/products")

	InitializeRoutes(engine)

	engine.Run()
}

func migrate() {
	database.Migrate()
}

