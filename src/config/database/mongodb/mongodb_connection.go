package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL     = "MONGODB_URL"
	MONGODB_USER_DB = "MONGODB_USER_DATABASE"
)

func NewMongodbConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URL)
	mongodb_database := os.Getenv(MONGODB_USER_DB)

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Println(err)
		return nil, err
	}

	return client.Database(mongodb_database), nil

}
