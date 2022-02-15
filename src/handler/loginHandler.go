package handler

import (
	. "blog-on-containers/constants"
	"blog-on-containers/entities"
	"blog-on-containers/models"
	"blog-on-containers/services"
	"blog-on-containers/token"
	"blog-on-containers/utils"
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
		badRequest(context, http.StatusBadRequest, MESSAGE_INVALID_REQUEST, err)
		return
	}

	userService := services.NewUserService()
	genrateJWTToken(context, loginObj, userService)
}

func genrateJWTToken(context *gin.Context, loginObj models.LoginRequest, userService services.UserService) {
	if !userService.IsValidUsernameAndPassword(loginObj) {
		badRequest(context, http.StatusBadRequest, MESSAGE_INVALID_USER, nil)
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
		badRequest(context, http.StatusBadRequest, MESSAGE_TOKEN_NOT_GENERATION, []models.ErrorDetail{
			{
				ErrorType:    models.ErrorTypeError,
				ErrorMessage: err.Error(),
			},
		})
		return
	}

	//utils.SetCookie(context, token)

	initSampleData(context)

	ok(context, http.StatusOK, MESSAGE_TOKEN_GENERATION, token)
}

func initSampleData(context *gin.Context) {
	userService := services.NewUserService()

	hasData := userService.IsValidUsernameAndPassword(models.LoginRequest{UserName: "admin", Password: "admin"})
	if hasData {
		return
	}

	err := userService.CreateUser(models.LoginRequest{UserName: "admin", Password: "admin"})
	if err != nil {
		return
	}

	storyService := services.NewStoryService(context)

	cu := utils.GetCurrentUser(context)
	storyService.CreateStory(entities.NewStory("Golang - Fabrika Tasarım Deseni", "Desing Patternlar, yazılım sürecinde sıkça karşılaşılan sorunların ortak çözüm yoludur. Buradaki sorunlar, runtime sırasında alınan hatalar değildir. Peki nedir? Hiç kodunuzu incelediğinizde “burada bir şey eksik ama” dediğiniz oluyor mu? İşte tam orada devreye Design Patternlar giriyor. Aynı kodun çok defa tekrar etmesi, connect işlemlerinde kendini tekrar eden yapılar v.b. Bu gibi tasarımsal sorunların çözümü Design Patternlardır. Kısacası geçmişten günümüzü deneme yanılma yoluyla ortaya çıkan, ortak sorunların çözümü için oluşturulmuş kalıplardır.", cu.ID))
}
