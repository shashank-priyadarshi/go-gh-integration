package mongo

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func setMongoConnection() *mongo.Client {
	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("MongoURI"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection succeeded:")
	return client
}

func writeDataToCollection(collection *mongo.Collection, gitHubData interface{}) error {
	fmt.Println("collection: ", collection)

	insertResult, err := collection.InsertOne(context.TODO(), gitHubData)
	if err != nil {
		fmt.Println("Error while Inserting github data to MongoDB!")
	} else {
		fmt.Println("GitHub data saved successfully! ", insertResult)
	}

	return err
}

func StartConnection(gitHubData interface{}) error {
	client := setMongoConnection()
	db = client.Database(os.Getenv("DB_NAME"))

	err := writeDataToCollection(db.Collection(os.Getenv("COLLECTION_NAME")), gitHubData)
	return err
}
