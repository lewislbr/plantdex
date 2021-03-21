package storage

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDatabase() (*mongo.Collection, error) {
	mongodbURI := os.Getenv("USERS_MONGODB_URI")
	databaseName := os.Getenv("USERS_DATABASE_NAME")
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongodbURI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Users database ready ✅")

	collectionName := os.Getenv("USERS_COLLECTION_NAME")

	return client.Database(databaseName).Collection(collectionName), nil
}

func DisconnectDatabase(ctx context.Context, db *mongo.Collection) error {
	fmt.Println("Disconnecting database...")

	err := db.Database().Client().Disconnect(ctx)
	if err != nil {
		return err
	}

	return nil
}
