package splitdatabase

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenDBConnection() (*mongo.Client, error) {
	database := os.Getenv("MONGODB_CONNECTION")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(database))

	if err != nil {
		return nil, err
	}

	return client, err
}
