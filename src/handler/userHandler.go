package handler

import (
	"blog-on-containers/models"
	"blog-on-containers/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	var loginObj models.LoginRequest
	if err := context.ShouldBindJSON(&loginObj); err != nil {
		var errors []models.ErrorDetail = make([]models.ErrorDetail, 0, 1)
		errors = append(errors, models.ErrorDetail{
			ErrorType:    models.ErrorTypeValidation,
			ErrorMessage: fmt.Sprintf("%v", err),
		})
		badRequest(context, http.StatusBadRequest, "invalid request", errors)
		return
	}

	if err := loginObj.IsValid(); err != nil {
		badRequest(context, http.StatusBadRequest, "invalid request", err)
		return
	}

	userService := services.NewUserService()
	_, err := userService.GetUserByUsername(loginObj.UserName)
	if err == nil {
		badRequest(context, http.StatusBadRequest, "Username already taken ", nil)
		return
	}

	if userService.CreateUser(loginObj) != nil {
		badRequest(context, http.StatusBadRequest, "User not created", nil)
		return
	}

	ok(context, http.StatusOK, "User created", nil)
}
