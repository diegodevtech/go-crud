package test

import (
	"context"
	"encoding/json"
	"io"
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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserUpdate(t *testing.T){
	recorder := httptest.NewRecorder()
	ctx := GetTestGinContext(recorder)
	id := primitive.NewObjectID()

	_, err := Database.Collection("test_collection").InsertOne(
		context.Background(),
		bson.M{"_id": id, "name": "old", "age": 20, "email": "test@test.com"},
	)
	if err != nil {
		t.Fatal(err)
		return
	}

	param := []gin.Param{
		{
			Key: "userId",
			Value: id.Hex(),
		},
	}

	userRequest := request.UserRequest{
		Name: "Updated",
		Age: 50,
	}

	b, _ := json.Marshal(userRequest)
	stringReader := io.NopCloser(strings.NewReader(string(b)))

	MakeRequest(ctx, param, url.Values{}, "PUT", stringReader)
	UserController.UpdateUser(ctx)

	assert.EqualValues(t, http.StatusOK, recorder.Code)

	userEntity := entity.UserEntity{}

	filter := bson.D{{Key:"_id", Value:id}}
	_ = Database.Collection("test_collection").FindOne(context.Background(), filter).Decode(&userEntity)

	assert.EqualValues(t, userRequest.Name, userEntity.Name)
	assert.EqualValues(t, userRequest.Age, userEntity.Age)

}