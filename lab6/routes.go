package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	r.GET("/", indexRoute)
}

func indexRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}
