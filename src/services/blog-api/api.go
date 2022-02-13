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

	validateTokenHandle := middleware.ValidateToken()
	authHandle := middleware.Authorization([]int{1})

	initBlogRouteMap(api, validateTokenHandle, authHandle)
	initUserRouteMap(api, validateTokenHandle, authHandle)

	r.Run(":5000")
}

func initBlogRouteMap(route *gin.RouterGroup, validateTokenHandle, authHandle gin.HandlerFunc) {
	blog := route.Group("/blog")

	blog.POST("/", handler.CreateStory, validateTokenHandle, authHandle)
	blog.POST("/:id", handler.UpdateStory, validateTokenHandle, authHandle)
	blog.DELETE("/:id", handler.DeleteStory, validateTokenHandle, authHandle)

	blog.GET("/", handler.GetStories)
}

func initUserRouteMap(route *gin.RouterGroup, validateTokenHandle, authHandle gin.HandlerFunc) {
	user := route.Group("/user", validateTokenHandle, authHandle)

	user.GET("/", func(c *gin.Context) {
		c.AbortWithStatusJSON(200, gin.H{
			"status": "ok",
		})
	})
}
