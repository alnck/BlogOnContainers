package repository

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMongoRepo interface {
	Find(selector map[string]interface{}, v interface{}) error
	FindOne(selector map[string]interface{}, v interface{}) error
	CountDocuments(selector map[string]interface{}) (int64, error)
	InsertOne(v interface{}) (*mongo.InsertOneResult, error)
	UpdateOne(filter, update map[string]interface{}) (*mongo.UpdateResult, error)
	DeleteOne(filter map[string]interface{}) (*mongo.DeleteResult, error)
}

var lock sync.Mutex

type MongoRepository struct {
	Collection *mongo.Collection
}

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

func GetMongoRepository(collectionName string) *MongoRepository {
	initMongoDBInstance()

	return &MongoRepository{Collection: mongoDBInstance.Collection(collectionName)}
}

func (repo *MongoRepository) Find(selector map[string]interface{}, v interface{}) error {
	cursor, err := repo.Collection.Find(context.Background(), selector)
	cursor.All(context.Background(), v)
	return err
}

func (repo *MongoRepository) FindOne(selector map[string]interface{}, v interface{}) error {
	return repo.Collection.FindOne(context.Background(), selector).Decode(v)
}

func (repo *MongoRepository) CountDocuments(selector map[string]interface{}) (int64, error) {
	return repo.Collection.CountDocuments(context.Background(), selector)
}

func (repo *MongoRepository) InsertOne(v interface{}) (*mongo.InsertOneResult, error) {
	return repo.Collection.InsertOne(context.Background(), v)
}

func (repo *MongoRepository) UpdateOne(filter, update map[string]interface{}) (*mongo.UpdateResult, error) {
	return repo.Collection.UpdateOne(
		context.Background(),
		filter,
		update,
	)
}

func (repo *MongoRepository) DeleteOne(filter map[string]interface{}) (*mongo.DeleteResult, error) {
	return repo.Collection.DeleteOne(context.Background(), filter)
}
