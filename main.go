package main

import (
	"context"
	"log"

	"github.com/diegodevtech/go-crud/src/configuration/database/mongodb"
	"github.com/diegodevtech/go-crud/src/configuration/logger"
	"github.com/diegodevtech/go-crud/src/controller"
	"github.com/diegodevtech/go-crud/src/controller/routes"
	"github.com/diegodevtech/go-crud/src/model/repository"
	"github.com/diegodevtech/go-crud/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	logger.Info("STARTING APPLICATION")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	
	if err != nil {
		log.Fatalf("Error trying to connect to database, error: %s\n", err.Error())
		return
	}

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)
	
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	err = router.Run(":8080")

	if err != nil {
		log.Fatal(err)
	}

}