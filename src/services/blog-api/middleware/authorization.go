package middleware

import (
	"blog-on-containers/constants/projectcontextkeys"
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
				ErrorMessage: "You are not authorized to access this path",
			},
		},
		Status:  http.StatusUnauthorized,
		Message: "Unauthorized access",
	})
}

func ValidateToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.Request.Header.Get("apikey")
		referer := context.Request.Header.Get("Referer")
		valid, claims := token.VerifyToken(tokenString, referer)
		if !valid {
			ReturnUnauthorized(context)
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

	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), projectcontextkeys.USER_CONTEXT_KEY, user))
}

func Authorization(validRoles []int) gin.HandlerFunc {
	return func(context *gin.Context) {

		if len(context.Keys) == 0 {
			ReturnUnauthorized(context)
		}

		rolesVal := context.Keys["Roles"]
		fmt.Println("roles", rolesVal)
		if rolesVal == nil {
			ReturnUnauthorized(context)
		}

		roles := rolesVal.([]int)
		validation := make(map[int]int)
		for _, val := range roles {
			validation[val] = 0
		}

		for _, val := range validRoles {
			if _, ok := validation[val]; !ok {
				ReturnUnauthorized(context)
			}
		}
	}
}
