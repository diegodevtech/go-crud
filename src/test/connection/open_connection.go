package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ory/dockertest/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenConnection() (close func()) {

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
		return
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mongo",
		Tag:        "latest",
	})

	if err != nil {
		log.Fatalf("Could not create mongo container: %s", err)
		return
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://127.0.0.1:%s", resource.GetPort("27017/tcp"))))

	if err != nil {
		log.Println("Error trying to open connection")
		return
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Println("Error trying to open connection")
		return
	}

	// variavel global nao instanciada
	// Collection = client.Database(DatabaseName).Collection(CollectionName)

	close = func() {
		resource.Close()
		if err != nil {
			log.Println("Error trying to open connection")
			return
		}
	}
	return
}
