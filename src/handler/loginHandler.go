package handler

import (
	"blog-on-containers/models"
	"blog-on-containers/services"
	"blog-on-containers/token"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginHandler(context *gin.Context) {
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
		badRequest(context, http.StatusBadRequest, "invalid request", err)
		return
	}

	userService := services.NewUserService()
	genrateJWTToken(context, loginObj, userService)
}

func genrateJWTToken(context *gin.Context, loginObj models.LoginRequest, userService services.UserService) {
	if !userService.IsValidUsernameAndPassword(loginObj) {
		badRequest(context, http.StatusBadRequest, "invalid user", nil)
		return
	}

	var claims = &models.JwtClaims{}
	claims.Username = loginObj.UserName
	claims.Roles = []int{1, 2, 3}
	claims.Audience = context.Request.Header.Get("Referer") // get it from Referer header

	var tokenCreationTime = time.Now().UTC()
	var expirationTime = tokenCreationTime.Add(time.Duration(2) * time.Hour)
	token, err := token.GenrateToken(claims, expirationTime)

	if err != nil {
		badRequest(context, http.StatusBadRequest, "error in gerating token", []models.ErrorDetail{
			{
				ErrorType:    models.ErrorTypeError,
				ErrorMessage: err.Error(),
			},
		})
		return
	}

	//utils.SetCookie(context, token)

	ok(context, http.StatusOK, "token created", token)
}
