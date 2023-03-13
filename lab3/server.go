package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Static("./static", "./static/")
	engine.LoadHTMLGlob("templates/*.tmpl")
	engine.Static("/image", "./resources")

	InitializeRoutes(engine)

	engine.Run()
}
