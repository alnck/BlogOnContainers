package handler

import (
	. "blog-on-containers/constants"
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
		badRequest(context, http.StatusBadRequest, MESSAGE_INVALID_REQUEST, err)
		return
	}

	cu := utils.GetCurrentUser(context)
	newStory := entities.NewStory(story.Title, story.Content, cu.ID)

	storyService := services.NewStoryService(context)
	if !storyService.CreateStory(newStory) {
		badRequest(context, http.StatusBadRequest, MESSAGE_STORY_NOT_CREATE, nil)
		return
	}

	ok(context, http.StatusCreated, MESSAGE_STORY_CREATE, story)
}

func UpdateStory(context *gin.Context) {
	var story models.StoryRequest

	if !shouldBindJSON(context, &story) {
		return
	}

	if err := story.IsValid(); err != nil {
		badRequest(context, http.StatusBadRequest, MESSAGE_INVALID_REQUEST, err)
		return
	}

	storyService := services.NewStoryService(context)
	if !storyService.UpdateStory(story) {
		badRequest(context, http.StatusBadRequest, MESSAGE_STORY_NOT_UPDATE, nil)
		return
	}

	ok(context, http.StatusCreated, MESSAGE_STORY_UPDATE, story)
}

func DeleteStory(context *gin.Context) {
	storyService := services.NewStoryService(context)
	if !storyService.DeleteStory() {
		badRequest(context, http.StatusBadRequest, MESSAGE_STORY_NOT_DELETE, nil)
		return
	}

	ok(context, http.StatusOK, MESSAGE_STORY_DELETE, nil)
}

func GetStories(context *gin.Context) {
	storyService := services.NewStoryService(context)
	stories := storyService.GetStories()

	ok(context, http.StatusOK, MESSAGE_STORIES_TAKEN, stories)
}

func GetStory(context *gin.Context) {
	storyService := services.NewStoryService(context)
	story, bool := storyService.GetStory()
	if !bool {
		badRequest(context, http.StatusBadRequest, MESSAGE_STORY_NOT_TAKEN, nil)
		return
	}

	ok(context, http.StatusOK, MESSAGE_STORY_TAKEN, story)
}
