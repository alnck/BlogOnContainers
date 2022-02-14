package main

import (
	"blog-on-containers/handler"
	"blog-on-containers/middleware"
	"blog-on-containers/models"
	"blog-on-containers/templates"
	"net/http"

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

	r.LoadHTMLGlob("templates/views/*.html")
	r.Static("/css", "./static/css")

	r.GET("/loginpage", templates.LoginPage)
	r.POST("/registerpage", templates.RegisterPage)

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

	user.GET("/", func(context *gin.Context) {
		context.AbortWithStatusJSON(http.StatusOK, models.Response{
			Data:    "ok",
			Status:  http.StatusOK,
			Message: "tes",
		})
	})

	user.POST("/", handler.CreateUser)
}
