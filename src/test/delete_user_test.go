package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDeleteUser(t *testing.T){

	recorder := httptest.NewRecorder()
	ctx := GetTestGinContext(recorder)
	id := primitive.NewObjectID()

	_, err := Database.Collection("test_collection").InsertOne(
		context.Background(),
		bson.M{"_id": id, "name": t.Name(), "email": "test@test.com"},
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

	MakeRequest(ctx, param, url.Values{}, "DELETE", nil)
	UserController.DeleteUser(ctx)

	assert.EqualValues(t, http.StatusOK, recorder.Code)

	filter := bson.D{{Key: "_id", Value: id}}
	result := Database.Collection("test_collection").FindOne(context.Background(),filter)

	assert.NotNil(t, result.Err())
}