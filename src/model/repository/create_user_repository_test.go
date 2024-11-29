package repository

import (
	"os"
	"testing"

	"github.com/diegodevtech/go-crud/src/model"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_CreateUser(t *testing.T){
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	os.Setenv("MONGODB_USER_DB", collectionName)
	defer os.Clearenv()

	mTestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	//  nao consegui fazer o mTestDb.CLose()

	mTestDb.Run("success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain(
			"a@a.com", "Teste@@2024", "Teste", 20,
		)
		userDomain, err := repo.CreateUser(domain)

		_, errId := primitive.ObjectIDFromHex(userDomain.GetID())

		assert.Nil(t, err)
		assert.Nil(t, errId)
		assert.EqualValues(t, userDomain.GetEmail(), domain.GetEmail())
		assert.EqualValues(t, userDomain.GetAge(), domain.GetAge())
		assert.EqualValues(t, userDomain.GetPassword(), domain.GetPassword())
		assert.EqualValues(t, userDomain.GetName(), domain.GetName())
	})

	mTestDb.Run("Error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{{Key:"ok",Value:0}})
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain(
			"a@a.com", "Teste@@2024", "Teste", 20,
		)
		userDomain, err := repo.CreateUser(domain)

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
	
}