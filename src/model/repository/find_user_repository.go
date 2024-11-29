package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/model"
	"github.com/diegodevtech/go-crud/src/model/repository/entity"
	"github.com/diegodevtech/go-crud/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Initializing FindUserByEmail Repository Method", zap.String("journey", "findUserByEmail"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email :%s", email)
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByEmail repository executed successfully", zap.String("journey", "findUserByEmail"), zap.String("email", email), zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Initializing FindUserByEmailAndPassword Repository Method", zap.String("journey", "findUserByEmailAndPassword"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "User not found with this email."
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmailAndPassword"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmailAndPassword"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(password))

	if err != nil {
		errorMessage := "Incorrect password"
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmailAndPassword"))
		return nil, rest_err.NewUnauthorizedError(errorMessage)
	}

	logger.Info("FindUserByEmailAndPassword repository executed successfully", zap.String("journey", "findUserByEmailAndPassword"), zap.String("email", email), zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Initializing FindUserByID Repository Method", zap.String("journey", "findUserByID"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this id: %s", id)
			logger.Error(errorMessage, err, zap.String("journey", "findUserByID"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by id"
		logger.Error(errorMessage, err, zap.String("journey", "findUserByID"))
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("FindUserByID repository executed successfully", zap.String("journey", "findUserByID"), zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}
