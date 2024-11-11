package mongodb

import (
	"context"

	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InitConnection() {
	ctx := context.Background()
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	logger.Info("Successfully connected to DB.")
}