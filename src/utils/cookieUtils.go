package utils

import (
	"github.com/gin-gonic/gin"
)

func SetCookie(context *gin.Context, token string) {
	context.SetCookie("token", token, 10, "/", "localhost", false, true)
}

func GetCookieValue(context *gin.Context) string {
	val, err := context.Cookie("token")
	if err != nil {
		return ""
	}

	return val
}

func RemoveCookie(context *gin.Context, username string) {
	context.SetCookie("token", "", -1, "/", "localhost", false, true)
}
