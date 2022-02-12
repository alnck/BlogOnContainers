package services

import (
	. "blog-on-containers/entities"
	"blog-on-containers/models"
	"blog-on-containers/repository"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	repo *repository.MongoRepository
)

const m_COLLECTION_NAME = "users"

type UserService struct{}

func NewUserService() UserService {
	repo = repository.GetMongoRepository(m_COLLECTION_NAME)

	return UserService{}
}

func (*UserService) IsValidUsernameAndPassword(loginObj models.LoginRequest) bool {
	filter := bson.M{"username": loginObj.UserName, "password": loginObj.Password}
	count, err := repo.CountDocuments(filter)

	return err == nil && count > 0
}

func (*UserService) GetUserByUsername(username string) (User, error) {
	var user User
	filter := bson.M{"username": username}
	err := repo.FindOne(filter, &user)

	return user, err
}
