package handler

import (
	"blog-on-containers/models"
	"blog-on-containers/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	var loginObj models.LoginRequest

	if !shouldBindJSON(context, &loginObj) {
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
