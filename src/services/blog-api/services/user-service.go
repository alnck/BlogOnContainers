package services

import (
	. "blog-on-containers/entities"
	"blog-on-containers/models"
	"blog-on-containers/repository"

	"go.mongodb.org/mongo-driver/bson"
)

type IUserService interface {
	IsValidUsernameAndPassword(loginObj models.LoginRequest) bool
	GetUserByUsername(username string) (User, error)
}

var (
	repoUsers repository.IMongoRepository
)

const m_COLLECTION_NAME_USERS = "users"

type UserService struct{}

func NewUserService() UserService {
	repoUsers = repository.NewMongoRepository()

	return UserService{}
}

func NewUserServiceForTest(mongoRepo repository.IMongoRepository) UserService {
	repoUsers = mongoRepo
	return UserService{}
}

func (*UserService) IsValidUsernameAndPassword(loginObj models.LoginRequest) bool {
	filter := bson.M{"username": loginObj.UserName, "password": loginObj.Password}
	count, err := repoUsers.CountDocuments(m_COLLECTION_NAME_USERS, filter)

	return err == nil && count > 0
}

func (*UserService) GetUserByUsername(username string) (User, error) {
	var user User
	filter := bson.M{"username": username}
	err := repoUsers.FindOne(m_COLLECTION_NAME_USERS, filter, &user)

	return user, err
}
