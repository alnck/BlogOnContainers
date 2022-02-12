package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Story struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CreatedDate  time.Time          `bson:"createddate,omitempty"`
	ModifiedDate time.Time          `bson:"modifieddate,omitempty"`
	DisplayRate  int                `bson:"displayrate,omitempty"`
	UserID       primitive.ObjectID `bson:"userid,omitempty"`
	Title        string             `bson:"title,omitempty"`
	Content      string             `bson:"content,omitempty"`
}

func NewStory(title, content string, userId primitive.ObjectID) Story {
	story := Story{
		CreatedDate: time.Now().UTC(),
		DisplayRate: 0,
		UserID:      userId,
		Title:       title,
		Content:     content,
	}

	return story
}
