package main

import (
	"blog-on-containers/handler"
	"blog-on-containers/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", handler.LoginHandler)

	api := r.Group("/api")

	api.Use(middleware.ValidateToken())

	product := api.Group("/product")

	product.Use(middleware.Authorization([]int{1}))

	user := api.Group("/User")
	user.GET("/", func(c *gin.Context) {
		c.AbortWithStatusJSON(200, gin.H{
			"status": "ok",
		})
	})
	r.Run(":5000")
}
