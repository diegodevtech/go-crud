package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/diegodevtech/go-crud/src/model/repository/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"golang.org/x/crypto/bcrypt"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv("MONGODB_USER_DB", databaseName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mTestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mTestDb.Run("success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "teste@teste.com",
			Password: "Teste@@2020",
			Name:     "Teste",
			Age:      20,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			ConvertEntityToBson(userEntity),
		))

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByEmail(userEntity.Email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
	})

	mTestDb.Run("error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("test")

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Error trying to find user by email")
		assert.Nil(t, userDomain)
	})

	mTestDb.Run("no document found error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
		))

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("test")

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "User not found with this email :test")
		assert.Nil(t, userDomain)
	})
}

func TestUserRepository_FindUserByEmailAndPassword(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mTestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mTestDb.Run("success", func(mt *mtest.T){
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("Teste@@2020"), bcrypt.DefaultCost)
		userEntity := entity.UserEntity{
			ID: primitive.NewObjectID(),
			Email: "teste@teste.com",
			Password: string(hashedPassword),
			Name: "Teste",
			Age: 20,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			ConvertEntityToBson(userEntity),
		))

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByEmailAndPassword(userEntity.Email, "Teste@@2020")

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
	})

	mTestDb.Run("user not found by email", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
		))
		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByEmailAndPassword("notfound@teste.com", "WrongPassword")

		assert.Nil(t, userDomain)
		assert.NotNil(t, err)
		assert.Equal(t, "User not found with this email.", err.Message)
	})

	mTestDb.Run("unexpected error in FindOne", func(mt *mtest.T) {

    mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
        Code:    11000,
        Message: "Simulated database error",
    }))

    databaseMock := mt.Client.Database(databaseName)
    repo := NewUserRepository(databaseMock)

    userDomain, err := repo.FindUserByEmailAndPassword("error@teste.com", "Password123")

    assert.Nil(t, userDomain)
    assert.NotNil(t, err)
    assert.Equal(t, "Error trying to find user by email", err.Message)
})

	mTestDb.Run("password incorrect", func(mt *mtest.T) {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("CorrectPassword@@2020"), bcrypt.DefaultCost)

		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "teste@teste.com",
			Password: string(hashedPassword),
			Name:     "Teste",
			Age:      20,
		}

		// Add mocked response
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			ConvertEntityToBson(userEntity),
		))

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)

		// Call with incorrect password
		userDomain, err := repo.FindUserByEmailAndPassword(userEntity.Email, "WrongPassword")

		// Assertions
		assert.Nil(t, userDomain)
		assert.NotNil(t, err)
		assert.Equal(t, "Incorrect password", err.Message)
	})
}

func TestUserRepository_FindUserByID(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mTestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mTestDb.Run("success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "teste@teste.com",
			Password: "Teste@@2020",
			Name:     "Teste",
			Age:      20,
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
			ConvertEntityToBson(userEntity),
		))

		databaseMock := mt.Client.Database(databaseName)
		repo := NewUserRepository(databaseMock)

		userDomain, err := repo.FindUserByID(userEntity.ID.Hex())

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
		assert.EqualValues(t, userDomain.GetName(), userEntity.Name)
		assert.EqualValues(t, userDomain.GetPassword(), userEntity.Password)
		assert.EqualValues(t, userDomain.GetAge(), userEntity.Age)
	})

	mTestDb.Run("error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByID("test")

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Error trying to find user by id")
		assert.Nil(t, userDomain)
	})

	mTestDb.Run("no document found error", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(
			0,
			fmt.Sprintf("%s.%s", databaseName, collectionName),
			mtest.FirstBatch,
		))

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByID("test")

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "User not found with this id: test")
		assert.Nil(t, userDomain)
	})
}

func ConvertEntityToBson(userEntity entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: userEntity.ID},
		{Key: "email", Value: userEntity.Email},
		{Key: "password", Value: userEntity.Password},
		{Key: "name", Value: userEntity.Name},
		{Key: "age", Value: userEntity.Age},
	}
}
