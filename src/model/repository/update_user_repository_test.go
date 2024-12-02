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

func TestUserRepository_UpdateUser(t *testing.T){
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mTestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mTestDb.Run("when_sending_a_valid_user_returns_success", func(mt *mtest.T) {
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
		err := repo.UpdateUser(domain.GetID(), domain)

		assert.Nil(t,err)
	})

	mTestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		domain := model.NewUserUpdateDomain("Novo nome", 20)
		domain.SetID(primitive.NewObjectID().Hex())
		err := repo.UpdateUser(domain.GetID(), domain)
		assert.NotNil(t, err)
	})
}