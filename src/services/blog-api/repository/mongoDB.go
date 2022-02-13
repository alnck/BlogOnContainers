package repository

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMongoRepository interface {
	Find(collectionName string, selector map[string]interface{}, v interface{}) error
	FindOne(collectionName string, selector map[string]interface{}, v interface{}) error
	CountDocuments(collectionName string, selector map[string]interface{}) (int64, error)
	InsertOne(collectionName string, v interface{}) (*mongo.InsertOneResult, error)
	UpdateOne(collectionName string, filter, update map[string]interface{}) (*mongo.UpdateResult, error)
	DeleteOne(collectionName string, filter map[string]interface{}) (*mongo.DeleteResult, error)
}

var lock sync.Mutex

type mongoRepository struct{}

var mongoDBInstance *mongo.Database

func initMongoDBInstance() {
	if mongoDBInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if mongoDBInstance == nil {
			clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") //mongodb://blogdb:27017
			client, err := mongo.Connect(context.TODO(), clientOptions)
			if err != nil {
				log.Fatal("⛒ Connection Failed to Database")
				log.Fatal(err)
			}

			// Check the connection
			err = client.Ping(context.TODO(), nil)
			if err != nil {
				log.Fatal("⛒ Connection Failed to Database")
				log.Fatal(err)
			}

			mongoDBInstance = client.Database("blogdb")
		}
	}
}

func NewMongoRepository() IMongoRepository {
	initMongoDBInstance()

	return &mongoRepository{}
}

func (*mongoRepository) Find(collectionName string, selector map[string]interface{}, v interface{}) error {
	cursor, err := mongoDBInstance.Collection(collectionName).Find(context.Background(), selector)
	cursor.All(context.Background(), v)
	return err
}

func (*mongoRepository) FindOne(collectionName string, selector map[string]interface{}, v interface{}) error {
	return mongoDBInstance.Collection(collectionName).FindOne(context.Background(), selector).Decode(v)
}

func (*mongoRepository) CountDocuments(collectionName string, selector map[string]interface{}) (int64, error) {
	return mongoDBInstance.Collection(collectionName).CountDocuments(context.Background(), selector)
}

func (*mongoRepository) InsertOne(collectionName string, v interface{}) (*mongo.InsertOneResult, error) {
	return mongoDBInstance.Collection(collectionName).InsertOne(context.Background(), v)
}

func (*mongoRepository) UpdateOne(collectionName string, filter, update map[string]interface{}) (*mongo.UpdateResult, error) {
	return mongoDBInstance.Collection(collectionName).UpdateOne(
		context.Background(),
		filter,
		update,
	)
}

func (*mongoRepository) DeleteOne(collectionName string, filter map[string]interface{}) (*mongo.DeleteResult, error) {
	return mongoDBInstance.Collection(collectionName).DeleteOne(context.Background(), filter)
}
