package main

import (
	"blog-on-containers/handler"
	"blog-on-containers/middleware"
	"blog-on-containers/templates"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge:   60 * 60 * 24, // expire in a day
		Path:     "/",
		Secure:   true,
		HttpOnly: true})

	r.Use(sessions.Sessions("mysession", store))

	r = gin.New()
	r.Use(gin.Logger())

	r.POST("/login", handler.LoginHandler)
	r.POST("/user", handler.CreateUser)

	api := r.Group("/api")

	validateTokenHandle := middleware.ValidateToken()
	authHandle := middleware.Authorization([]int{1})

	initTemplatesRouteMap(r)
	initBlogRouteMap(api, validateTokenHandle, authHandle)

	r.Run(":5000")
}

func initTemplatesRouteMap(route *gin.Engine) {
	route.Use(gin.Logger())

	route.LoadHTMLGlob("templates/views/*.html")
	route.Static("/css", "./static/css")

	route.GET("/loginpage", templates.LoginPage)
	route.GET("/registerpage", templates.RegisterPage)
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
