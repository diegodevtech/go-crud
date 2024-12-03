package service

import (
	"os"
	"testing"

	"github.com/diegodevtech/go-crud/src/configuration/rest_err"
	"github.com/diegodevtech/go-crud/src/model"
	"github.com/diegodevtech/go-crud/src/test/mocks"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_calling_login_repository_returns_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("teste@teste.com", "Teste@@2020", "Teste", 20)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmailAndPassword(
			userDomain.GetEmail(),
			userDomain.GetPassword(),
		).Return(nil, rest_err.NewInternalServerError("Error trying to find user by email and password."))

		user, token, err := service.LoginService(userDomain)

		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Error trying to find user by email and password.")
	})

	t.Run("when_generating_token_returns_error", func(t *testing.T) {
		userDomainMock := mocks.NewMockUserDomainInterface(ctrl)

		userDomainMock.EXPECT().GetEmail().Return("teste@teste.com")
		userDomainMock.EXPECT().GetPassword().Return("Teste@@2020")
		userDomainMock.EXPECT().GenerateToken().Return("", rest_err.NewInternalServerError("Error trying to generate token."))
		
		repository.EXPECT().FindUserByEmailAndPassword("teste@teste.com", "Teste@@2020").Return(userDomainMock, nil)

		user, token, err := service.LoginService(userDomainMock)

		assert.Nil(t, user)
		assert.Empty(t, token)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Error trying to generate token.")
	})

	t.Run("when_user_and_password_is_valid_returns_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		secret := "test"
		os.Setenv("JWT_SECRET_KEY", secret)
		defer os.Clearenv()

		userDomain := model.NewUserDomain("teste@teste.com", "Teste@@2020", "Teste", 20)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmailAndPassword(
			userDomain.GetEmail(),
			gomock.Any(),
		).Return(userDomain, nil)

		user, token, err := service.LoginService(userDomain)

		assert.NotNil(t, user)
		assert.Nil(t, err)
		assert.EqualValues(t, user.GetID(), id)
		assert.EqualValues(t, user.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, user.GetName(), userDomain.GetName())
		assert.EqualValues(t, user.GetPassword(), userDomain.GetPassword())
		assert.EqualValues(t, user.GetAge(), userDomain.GetAge())

		tokenReturned, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
				return []byte(secret), nil
			}
			return nil, rest_err.NewBadRequestError("Invalid token")
		})
		_, ok := tokenReturned.Claims.(jwt.MapClaims)
		if !ok || !tokenReturned.Valid {
			t.FailNow()
			return
		}
	})
}
