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
	initUserRouteMap(api)

	r.Run(":5000")
}

func initBlogRouteMap(route *gin.RouterGroup, validateTokenHandle, authHandle gin.HandlerFunc) {
	blog := route.Group("/blog")
	blog.Use(validateTokenHandle, authHandle)

	blog.POST("/", handler.CreateStory)
	blog.POST("/:id", handler.UpdateStory)
	blog.DELETE("/:id", handler.DeleteStory)

	blogWithoutAuth := route.Group("/blog")

	blogWithoutAuth.GET("/", handler.GetStories)
}

func initUserRouteMap(route *gin.RouterGroup) {
	user := route.Group("/user")

	user.POST("/", handler.CreateUser)
}
