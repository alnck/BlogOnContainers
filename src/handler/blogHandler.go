package handler

import (
	"blog-on-containers/entities"
	"blog-on-containers/models"
	"blog-on-containers/services"
	"blog-on-containers/utils"
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

	if err := story.IsValid(); err != nil {
		badRequest(context, http.StatusBadRequest, "invalid request", err)
		return
	}

	cu := utils.GetCurrentUser(context)
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

	if err := story.IsValid(); err != nil {
		badRequest(context, http.StatusBadRequest, "invalid request", err)
		return
	}

	storyService := services.NewStoryService(context)
	storyService.UpdateStory(story)

}

func DeleteStory(context *gin.Context) {
	storyService := services.NewStoryService(context)
	storyService.DeleteStory()

}

func GetStories(context *gin.Context) {
	storyService := services.NewStoryService(context)
	stories := storyService.GetStories()

	ok(context, http.StatusOK, "All Stories Taken", stories)
}
