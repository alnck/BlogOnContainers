package templates

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPage(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", nil)
}

func RegisterPage(context *gin.Context) {
	context.HTML(http.StatusOK, "register.html", nil)
}
