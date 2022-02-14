package utils

import (
	. "blog-on-containers/constants/ProjectContextKeys"
	. "blog-on-containers/entities"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(context *gin.Context) User {
	return context.Request.Context().Value(USER_CONTEXT_KEY).(User)
}
