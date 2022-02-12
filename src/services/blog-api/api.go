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

	initBlogRouteMap(api)
	initUserRouteMap(api)

	r.Run(":5000")
}

func initBlogRouteMap(route *gin.RouterGroup) {
	blog := route.Group("/blog")
	blog.Use(middleware.Authorization([]int{1}))

	blog.POST("/", handler.CreateStory)
	blog.POST("/:id", handler.UpdateStory)

}

func initUserRouteMap(route *gin.RouterGroup) {
	user := route.Group("/user")
	user.GET("/", func(c *gin.Context) {
		c.AbortWithStatusJSON(200, gin.H{
			"status": "ok",
		})
	})
}
