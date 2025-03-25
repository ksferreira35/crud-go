package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	mongodb "github.com/ksferreira35/crud-go/src/config/database/mongodb"
	"github.com/ksferreira35/crud-go/src/config/logger"
	"github.com/ksferreira35/crud-go/src/controller"
	"github.com/ksferreira35/crud-go/src/controller/routes"
	"github.com/ksferreira35/crud-go/src/model/repository"
	"github.com/ksferreira35/crud-go/src/model/service"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	database, err := mongodb.NewMongodbConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", 
		err.Error())
		return
	}

	// Init dependecies
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
