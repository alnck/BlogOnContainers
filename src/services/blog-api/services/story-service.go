package services

import (
	"blog-on-containers/entities"
	"blog-on-containers/helper"
	"blog-on-containers/models"
	"blog-on-containers/repository"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IStoryService interface {
	CreateStory(story entities.Story)
	UpdateStory(story models.StoryRequest) bool
	DeleteStory() bool
	GetStories() ([]entities.Story, error)
}

var (
	repoStories  *repository.MongoRepository
	storyContext *gin.Context
)

const m_COLLECTION_NAME_STORIES = "stories"

type StoryService struct{}

func NewStoryService(context *gin.Context) StoryService {
	repoStories = repository.GetMongoRepository(m_COLLECTION_NAME_STORIES)
	storyContext = context

	return StoryService{}
}

func (*StoryService) CreateStory(story entities.Story) {
	repoStories.InsertOne(story)
}

func (*StoryService) UpdateStory(story models.StoryRequest) bool {
	storyId := storyContext.Param("id")
	id, err := primitive.ObjectIDFromHex(storyId)
	if err != nil {
		return false
	}

	cu := helper.GetCurrentUser(storyContext)

	filter := bson.M{"_id": id, "userid": cu.ID}
	update := bson.M{
		"$set": bson.M{
			"title":        story.Title,
			"content":      story.Content,
			"modifieddate": time.Now().UTC(),
		},
	}

	repoStories.UpdateOne(filter, update)

	return true
}

func (*StoryService) DeleteStory() bool {
	storyId := storyContext.Param("id")
	id, err := primitive.ObjectIDFromHex(storyId)
	if err != nil {
		return false
	}

	cu := helper.GetCurrentUser(storyContext)
	filter := bson.M{"_id": id, "userid": cu.ID}

	repoStories.DeleteOne(filter)

	return true

}

func (*StoryService) GetStories() ([]entities.Story, error) {
	var stories []entities.Story
	filter := bson.M{}

	err := repoStories.Find(filter, &stories)

	return stories, err
}
