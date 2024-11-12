package repository

import (
	"context"
	"os"

	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/model"
	"github.com/diegodevtech/go-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Initializing CreateUser Repository Method", zap.String("journey", "createUser"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Eror trying to create user", err, zap.String("journey","createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(bson.ObjectID)

	logger.Info("CreateUser repository executed successfully",
		zap.String("userId", value.ID.Hex()),
		zap.String("journey", "createUser"),
	)

	return converter.ConvertEntityToDomain(*value), nil
}