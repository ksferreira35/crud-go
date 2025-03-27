package main

import (
	"github.com/ksferreira35/crud-go/src/controller"
	"github.com/ksferreira35/crud-go/src/model/repository"
	"github.com/ksferreira35/crud-go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	// Init dependecies
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}