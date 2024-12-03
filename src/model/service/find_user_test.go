package service

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/model"
	"github.com/diegodevtech/go-crud/src/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserService_FindUserByID(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("teste@teste.com", "Teste@@2020", "Teste", 20)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByID(id).Return(userDomain, nil)

		userResult, err := service.FindUserByIDService(id)

		assert.Nil(t, err)
		assert.NotNil(t, userResult)
		assert.EqualValues(t, userResult.GetID(), id)
		assert.EqualValues(t, userResult.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userResult.GetName(), userDomain.GetName())
		assert.EqualValues(t, userResult.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userResult.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		
		repository.EXPECT().FindUserByID(id).Return(nil, rest_err.NewNotFoundError("User Not Found"))
		
		userResult, err := service.FindUserByIDService(id)

		assert.Nil(t, userResult)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "User Not Found")
	})
}

func TestUserService_FindUserByEmail(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "teste@teste.com"
		userDomain := model.NewUserDomain(email, "Teste@@2020", "Teste", 20)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(email).Return(userDomain, nil)

		userResult, err := service.FindUserByEmailService(email)

		assert.Nil(t, err)
		assert.NotNil(t, userResult)
		assert.EqualValues(t, userResult.GetID(), id)
		assert.EqualValues(t, userResult.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userResult.GetName(), userDomain.GetName())
		assert.EqualValues(t, userResult.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userResult.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		email:= "teste@teste.com"

		repository.EXPECT().FindUserByEmail(email).Return(nil, rest_err.NewNotFoundError("User Not Found"))
		
		userResult, err := service.FindUserByEmailService(email)

		assert.Nil(t, userResult)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "User Not Found")
	})
}

func TestUserService_FindUserByEmailAndPassword(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	// service := NewUserDomainService(repository) isso foi feito pois findUserByEMailAndPassword está privado
	service := &userDomainService{repository}

	t.Run("when_exists_an_user_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "teste@teste.com"
		password := strconv.FormatInt(rand.Int63(), 10)

		userDomain := model.NewUserDomain(email, password, "Teste", 20)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(userDomain, nil)

		userResult, err := service.findUserByEmailAndPasswordService(email, password)

		assert.Nil(t, err)
		assert.NotNil(t, userResult)
		assert.EqualValues(t, userResult.GetID(), id)
		assert.EqualValues(t, userResult.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, userResult.GetName(), userDomain.GetName())
		assert.EqualValues(t, userResult.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, userResult.GetAge(), userDomain.GetAge())
	})

	t.Run("when_does_not_exists_an_user_returns_error", func(t *testing.T) {
		email := "teste@teste.com"
		password := strconv.FormatInt(rand.Int63(), 10)

		repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(nil, rest_err.NewNotFoundError("User Not Found"))
		
		userResult, err := service.findUserByEmailAndPasswordService(email, password)

		assert.Nil(t, userResult)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "User Not Found")
	})
}