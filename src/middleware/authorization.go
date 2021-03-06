package middleware

import (
	. "blog-on-containers/constants"
	"blog-on-containers/constants/ProjectContextKeys"
	"blog-on-containers/models"
	. "blog-on-containers/services"
	"blog-on-containers/token"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnUnauthorized(context *gin.Context) {
	context.AbortWithStatusJSON(http.StatusUnauthorized, models.Response{
		Error: []models.ErrorDetail{
			{
				ErrorType:    models.ErrorTypeUnauthorized,
				ErrorMessage: ERROR_MESSAGE_UNAUTHORIZED,
			},
		},
		Status:  http.StatusUnauthorized,
		Message: MESSAGE_UNAUTHORIZED,
	})
}

func ValidateToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.Request.Header.Get("apikey")
		/*tokenString := utils.GetCookieValue(context)
		if tokenString == "" {
			ReturnUnauthorized(context)
			return
		}*/
		referer := context.Request.Header.Get("Referer")
		valid, claims := token.VerifyToken(tokenString, referer)
		if !valid {
			ReturnUnauthorized(context)
			return
		}
		if len(context.Keys) == 0 {
			context.Keys = make(map[string]interface{})
		}
		context.Keys["Username"] = claims.Username
		context.Keys["Roles"] = claims.Roles

		addToContext(context, claims.Username)
	}
}

func addToContext(c *gin.Context, username string) {
	userService := NewUserService()
	user, err := userService.GetUserByUsername(username)

	if err != nil {
		ReturnUnauthorized(c)
	}

	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), ProjectContextKeys.USER_CONTEXT_KEY, user))
}

func Authorization(validRoles []int) gin.HandlerFunc {
	return func(context *gin.Context) {

		if len(context.Keys) == 0 {
			ReturnUnauthorized(context)
			return
		}

		rolesVal := context.Keys["Roles"]
		fmt.Println("roles", rolesVal)
		if rolesVal == nil {
			ReturnUnauthorized(context)
			return
		}

		roles := rolesVal.([]int)
		validation := make(map[int]int)
		for _, val := range roles {
			validation[val] = 0
		}

		for _, val := range validRoles {
			if _, ok := validation[val]; !ok {
				ReturnUnauthorized(context)
				return
			}
		}
	}
}
