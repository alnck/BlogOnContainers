package templates

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

var tmpl *template.Template

type todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []todo
}

func LoginPage(context *gin.Context) {

	context.HTML(http.StatusOK, "login.html", nil)

}
func RegisterPage(context *gin.Context) {

	context.HTML(http.StatusOK, "register.html", nil)

}
