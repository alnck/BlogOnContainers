package handler

import (
	. "blog-on-containers/constants"
	"blog-on-containers/models"
	"blog-on-containers/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(context *gin.Context) {
	var loginObj models.LoginRequest

	context.MultipartForm()
	for key, value := range context.Request.PostForm {
		if key == "Username" {
			loginObj.UserName = value[0]
		} else if key == "password" {
			loginObj.Password = value[0]
		}
	}

	if loginObj.UserName == "" && loginObj.Password == "" {
		if !shouldBindJSON(context, &loginObj) {
			return
		}
	}

	if err := loginObj.IsValid(); err != nil {
		badRequest(context, http.StatusBadRequest, MESSAGE_INVALID_REQUEST, err)
		return
	}

	userService := services.NewUserService()
	_, err := userService.GetUserByUsername(loginObj.UserName)
	if err == nil {
		badRequest(context, http.StatusBadRequest, MESSAGE_USER_ALREADY_TAKEN, nil)
		return
	}

	if userService.CreateUser(loginObj) != nil {
		badRequest(context, http.StatusBadRequest, MESSAGE_USER_NOT_CREATED, nil)
		return
	}

	ok(context, http.StatusOK, MESSAGE_USER_CREATE, nil)
}
