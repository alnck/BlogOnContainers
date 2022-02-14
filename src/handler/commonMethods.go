package handler

import (
	"blog-on-containers/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ok(context *gin.Context, status int, message string, data interface{}) {
	context.AbortWithStatusJSON(http.StatusOK, models.Response{
		Data:    data,
		Status:  status,
		Message: message,
	})
}
func badRequest(context *gin.Context, status int, message string, errors []models.ErrorDetail) {
	context.AbortWithStatusJSON(http.StatusBadRequest, models.Response{
		Error:   errors,
		Status:  status,
		Message: message,
	})

}

func shouldBindJSON(context *gin.Context, v interface{}) bool {
	if err := context.ShouldBindJSON(&v); err != nil {
		var errors []models.ErrorDetail = make([]models.ErrorDetail, 0, 1)
		errors = append(errors, models.ErrorDetail{
			ErrorType:    models.ErrorTypeValidation,
			ErrorMessage: fmt.Sprintf("%v", err),
		})
		badRequest(context, http.StatusBadRequest, "invalid request", errors)
		return false
	}

	return true
}
