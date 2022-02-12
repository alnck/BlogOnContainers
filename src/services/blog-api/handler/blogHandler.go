package handler

import (
	"blog-on-containers/entities"
	"blog-on-containers/helper"
	"blog-on-containers/models"
	"blog-on-containers/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStory(context *gin.Context) {
	var story models.StoryRequest
	if err := context.ShouldBindJSON(&story); err != nil {
		var errors []models.ErrorDetail = make([]models.ErrorDetail, 0, 1)
		errors = append(errors, models.ErrorDetail{
			ErrorType:    models.ErrorTypeValidation,
			ErrorMessage: fmt.Sprintf("%v", err),
		})
		badRequest(context, http.StatusBadRequest, "invalid request", errors)
	}

	cu := helper.GetCurrentUser(context)

	newStory := entities.NewStory(story.Title, story.Content, cu.ID)

	storyService := services.NewStoryService(context)
	storyService.CreateStory(newStory)

	ok(context, http.StatusCreated, "story Added", story)
}

func UpdateStory(context *gin.Context) {
	var story models.StoryRequest
	if err := context.ShouldBindJSON(&story); err != nil {
		var errors []models.ErrorDetail = make([]models.ErrorDetail, 0, 1)
		errors = append(errors, models.ErrorDetail{
			ErrorType:    models.ErrorTypeValidation,
			ErrorMessage: fmt.Sprintf("%v", err),
		})
		badRequest(context, http.StatusBadRequest, "invalid request", errors)
	}

	storyService := services.NewStoryService(context)

	storyService.UpdateStory(story)

}

func DeleteStory(context *gin.Context) {

	storyService := services.NewStoryService(context)

	storyService.DeleteStory()

}
