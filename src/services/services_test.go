package services

import (
	"blog-on-containers/constants/ProjectContextKeys"
	"blog-on-containers/entities"
	"blog-on-containers/models"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var invalidloginRequestList = []struct {
	UserName   string `json:"UserName" form:"UserName" binding:"required"`
	Password   string `json:"Password" form:"Password" binding:"required"`
	RememberMe bool   `json:"RememberMe" form:"RememberMe"`
}{
	{"", "", false},
	{"admin", "", false},
	{"", "1234", false},
	{"admin", "1234", false},
}

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) CountDocuments(collectionName string, selector map[string]interface{}) (int64, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(int64), args.Error(1)
}

func (mock *MockRepository) Find(collectionName string, selector map[string]interface{}, v interface{}) error {
	args := mock.Called(collectionName, selector, v)
	return args.Error(0)
}

func (mock *MockRepository) FindOne(collectionName string, selector map[string]interface{}, v interface{}) error {
	args := mock.Called()
	return args.Error(1)
}

func (mock *MockRepository) InsertOne(collectionName string, v interface{}) (*mongo.InsertOneResult, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*mongo.InsertOneResult), args.Error(1)
}

func (mock *MockRepository) UpdateOne(collectionName string, filter, update map[string]interface{}) (*mongo.UpdateResult, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*mongo.UpdateResult), args.Error(0)
}

func (mock *MockRepository) DeleteOne(collectionName string, filter map[string]interface{}) (*mongo.DeleteResult, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*mongo.DeleteResult), args.Error(1)
}

func createTestModeContext(request *http.Request) *gin.Context {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	objid, _ := primitive.ObjectIDFromHex("507f191e810c19729de860ea")
	tempUser := entities.User{ID: objid, Username: "test", Password: "test"}

	c.Request = request
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), ProjectContextKeys.USER_CONTEXT_KEY, tempUser))

	return c
}

func TestInValidUserNameAndPassword(t *testing.T) {
	mockRepo := new(MockRepository)

	//setup expectatinos
	mockRepo.On("CountDocuments").Return(int64(0), nil)

	testService := NewUserServiceForTest(mockRepo)

	for _, _data := range invalidloginRequestList {
		result := testService.IsValidUsernameAndPassword(_data)
		assert.False(t, result)
	}
}

func TestIfThereIsNoStory(t *testing.T) {
	mockRepo := new(MockRepository)

	//setup expectatinos
	mockRepo.On("Find").Return(errors.New("No data"))

	hr, _ := http.NewRequest("GET", "/", nil)
	testService := NewStoryServiceForTest(mockRepo, createTestModeContext(hr))

	stories := testService.GetStories()

	assert.Empty(t, stories)
}

/*func TestHasStoryDocuments(t *testing.T) {
	mockRepo := new(MockRepository)

	userid, _ := primitive.ObjectIDFromHex("507f191e810c19729de860ea")
	tempStories := []entities.Story{entities.NewStory("titleTest", "contentTest", userid)}

	//setup expectatinos
	mockRepo.On("Find", "stories", nil, tempStories).Return(nil)

	testService := NewStoryServiceForTest(mockRepo, createTestModeContext())

	stories := testService.GetStories()

	assert.True(t, len(stories) > 0)
}*/

/*func Test(t *testing.T) {
	mockRepo := new(MockRepository)
	var storyRequest models.StoryRequest

	//setup expectatinos
	mockRepo.On("InsertOne").Return(errors.New("Stories Not Have Tittle And Content"))

	context := createTestModeContext()
	context.Params = []gin.Param{gin.Param{Key: "id", Value: "507f191e810c19729de860ea"}}

	testService := NewStoryServiceForTest(mockRepo, context)

	err := testService.UpdateStory(storyRequest)
	assert.False(t, err)
}*/

func TestUpdateStoryOK(t *testing.T) {
	mockRepo := new(MockRepository)

	//setup expectatinos
	mockRepo.On("UpdateOne").Return(mongo.UpdateResult{UpsertedCount: 1}, true)

	hr, _ := http.NewRequest("GET", "/507f191e810c19729de860ea", nil)
	context := createTestModeContext(hr)
	context.Params = []gin.Param{gin.Param{Key: "id", Value: "507f191e810c19729de860ea"}}

	testService := NewStoryServiceForTest(mockRepo, context)

	result := testService.UpdateStory(models.StoryRequest{Title: "titleTest", Content: "content"})

	assert.True(t, result)
}
