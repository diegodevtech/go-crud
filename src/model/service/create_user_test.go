package service

import (
	"testing"

	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/model"
	"github.com/diegodevtech/go-crud/src/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserService_CreateUser(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_user_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("teste@teste.com", "Teste@@2020", "Teste", 20)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(userDomain, nil)

		user, err := service.CreateUserService(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Email is already registered in another account.")
	})

	t.Run("when_user_is_not_registered_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("teste@teste.com", "Teste@@2020", "Teste", 20)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)

		repository.EXPECT().CreateUser(userDomain).Return(nil, rest_err.NewInternalServerError("Error trying to create user"))

		user, err := service.CreateUserService(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Error trying to create user")
	})

	t.Run("when_user_is_not_registered_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("teste@teste.com", "Teste@@2020", "Teste", 20)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)

		repository.EXPECT().CreateUser(userDomain).Return(userDomain, nil)

		user, err := service.CreateUserService(userDomain)

		assert.EqualValues(t, user.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, user.GetName(), userDomain.GetName())
		assert.EqualValues(t, user.GetName(), userDomain.GetName())
		assert.EqualValues(t, user.GetAge(), userDomain.GetAge())
		assert.EqualValues(t, user.GetPassword(), userDomain.GetPassword())
		assert.NotNil(t, user)
		assert.Nil(t, err)
	})

}