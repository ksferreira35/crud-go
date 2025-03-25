package repository

import (
	"context"
	"os"

	"github.com/ksferreira35/crud-go/src/config/logger"
	rest_err "github.com/ksferreira35/crud-go/src/config/res_err"
	"github.com/ksferreira35/crud-go/src/model"
	"github.com/ksferreira35/crud-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createUser repository")
	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())	 
	}

	value.ID = result.InsertedID.(primitive.ObjectID)


	return converter.ConvertEntityToDomain(*value), nil 
}
