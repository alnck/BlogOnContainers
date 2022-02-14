package handler

import (
	"blog-on-containers/entities"
	"blog-on-containers/models"
	"blog-on-containers/services"
	"blog-on-containers/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStory(context *gin.Context) {
	var story models.StoryRequest

	if !shouldBindJSON(context, &story) {
		return
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

	if !shouldBindJSON(context, &story) {
		return
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
