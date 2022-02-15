package main

import (
	. "blog-on-containers/constants"
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

	r.POST(LINK_LOGIN_RELATIVEPATH, handler.LoginHandler)
	r.POST(LINK_USER_RELATIVEPATH, handler.CreateUser)

	api := r.Group(LINK_API_GROUP_RELATIVEPATH)

	validateTokenHandle := middleware.ValidateToken()
	authHandle := middleware.Authorization([]int{1})

	initTemplatesRouteMap(r)
	initBlogRouteMap(api, validateTokenHandle, authHandle)

	r.Run(SERVER_PORT)
}

func initTemplatesRouteMap(route *gin.Engine) {
	route.Use(gin.Logger())

	route.LoadHTMLGlob("templates/views/*.html")
	route.Static("/css", "./static/css")

	route.GET(LINK_TEMPLATE_LOGINPAGE_RELATIVEPATH, templates.LoginPage)
	route.GET(LINK_TEMPLATE_REGISTERPAGE_RELATIVEPATH, templates.RegisterPage)
	route.GET("/storypage/:id", templates.StoryPage)
}

func initBlogRouteMap(route *gin.RouterGroup, validateTokenHandle, authHandle gin.HandlerFunc) {
	blog := route.Group(LINK_BLOG_GROUP_RELATIVEPATH)
	blog.Use(validateTokenHandle, authHandle)

	blog.POST(LINK_RELATIVEPATH_CONSTANT, handler.CreateStory)
	blog.POST(LINK_API_BLOG_GROUP_ID_RELATIVEPATH, handler.UpdateStory)
	blog.DELETE(LINK_API_BLOG_GROUP_ID_RELATIVEPATH, handler.DeleteStory)

	blogWithoutAuth := route.Group(LINK_BLOG_GROUP_RELATIVEPATH)

	blogWithoutAuth.GET(LINK_RELATIVEPATH_CONSTANT, handler.GetStories)
	blogWithoutAuth.GET(LINK_API_BLOG_GROUP_ID_RELATIVEPATH, handler.GetStory)
}
