package repository

import (
	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/model"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/v2/mongo" //	tire o /v2 pcausa do teste
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(userId string) *rest_err.RestErr
}