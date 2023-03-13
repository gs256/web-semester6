package main

import (
	"lab4/application"
	"lab4/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// migrate()
	// panic("")
	engine := gin.Default()
	application.New().Run(engine)

	engine.Static("./static", "./static/")
	engine.LoadHTMLGlob("templates/**/*")
	engine.Static("/image", "./resources")

	InitializeRoutes(engine)

	engine.Run()
}

func migrate() {
	database.Migrate()
}
