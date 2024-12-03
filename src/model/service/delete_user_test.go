package service

import (
	"testing"

	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserService_DeleteUser(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_sending_a_valid_userId_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repository.EXPECT().DeleteUser(id).Return(nil)

		err := service.DeleteUserService(id)

		assert.Nil(t, err)
	})

	t.Run("when_sending_an_invalid_userId_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repository.EXPECT().DeleteUser(id).Return(rest_err.NewInternalServerError("Error trying to delete user."))

		err := service.DeleteUserService(id)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Error trying to delete user.")
	})
}