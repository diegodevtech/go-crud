package entity

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserEntity struct {
	ID       bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Email    string        `bson:"email,omitempty"`
	Password string        `bson:"password,omitempty"`
	Name     string        `bson:"name,omitempty"`
	Age      int8          `bson:"age,omitempty"`
}
