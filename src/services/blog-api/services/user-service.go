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
	repoUsers *repository.MongoRepository
)

const m_COLLECTION_NAME_USERS = "users"

type UserService struct{}

func NewUserService() UserService {
	repoUsers = repository.GetMongoRepository(m_COLLECTION_NAME_USERS)

	return UserService{}
}

func (*UserService) IsValidUsernameAndPassword(loginObj models.LoginRequest) bool {
	filter := bson.M{"username": loginObj.UserName, "password": loginObj.Password}
	count, err := repoUsers.CountDocuments(filter)

	return err == nil && count > 0
}

func (*UserService) GetUserByUsername(username string) (User, error) {
	var user User
	filter := bson.M{"username": username}
	err := repoUsers.FindOne(filter, &user)

	return user, err
}
