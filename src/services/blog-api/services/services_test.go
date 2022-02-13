package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	args := mock.Called()
	return args.Error(1)
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
	return result.(*mongo.UpdateResult), args.Error(1)
}

func (mock *MockRepository) DeleteOne(collectionName string, filter map[string]interface{}) (*mongo.DeleteResult, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*mongo.DeleteResult), args.Error(1)
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
