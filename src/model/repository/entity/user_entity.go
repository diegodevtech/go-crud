package entity

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UserEntity struct {
	ID       bson.ObjectID `bson:"_id,omitempty"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
	Name     string        `bson:"name"`
	Age      int8          `bson:"age"`
}
