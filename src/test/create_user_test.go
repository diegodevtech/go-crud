package test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/diegodevtech/go-crud/src/controller/model/request"
	"github.com/diegodevtech/go-crud/src/model/repository/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser(t *testing.T) {

	t.Run("user_already_registered_with_this_email", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		email := fmt.Sprintf("%d@test.com", rand.Int())

		_, err := Database.Collection("test_collection").InsertOne(context.Background(), bson.M{
			"name": t.Name(),
			"email": email,
		})

		if err != nil {
			t.Fatal(err)
			return
		}

		userRequest := request.UserRequest{
			Email: email,
			Password: "Bla@@@Bla12",
			Name: "Teste",
			Age: 20,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.CreateUser(ctx)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("user_is_not_registered_in_database", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		email := fmt.Sprintf("%d@test.com", rand.Int())

		userRequest := request.UserRequest{
			Email:    email,
			Password: "Teste@#@123",
			Name:     "Test User",
			Age:      50,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.CreateUser(ctx)

		userEntity := entity.UserEntity{}

		filter := bson.D{{Key: "email", Value: email}}
		_ = Database.
			Collection("test_collection").
			FindOne(context.Background(), filter).Decode(&userEntity)

		err := bcrypt.CompareHashAndPassword([]byte(userEntity.Password), []byte(userRequest.Password))

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, userEntity.Email, userRequest.Email)
		assert.EqualValues(t, userEntity.Name, userRequest.Name)
		assert.EqualValues(t, userEntity.Age, userRequest.Age)
		assert.Nil(t, err)
		// assert.EqualValues(t, userEntity.Password, userRequest.Password)
	})

}
