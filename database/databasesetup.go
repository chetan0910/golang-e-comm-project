package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/x/mongo/driver/mongocrypt/options"
)

func DBSet() *mongo.Client {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("failed to connect tomongodb :(")
		return nil
	}

	fmt.Println("successfully connected to mongodb")
	return client

}

var client *mongo.Client = DBSet() //whatever the DBSet() returning we are taking it in variable client.

func UserData(client *mongo.Client, collectionNAME string) *mongo.Collection {

}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {

}
