package dataBase

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var conn = InitMongoDB()
var db = DbConnect()

func DbConnect() *mongo.Database {
	return conn.Database("myfiles")
}
func InitCollections() (*mongo.Collection, *mongo.Collection, context.Context) {
	fsFiles := db.Collection("fs.files")
	fsTokens := db.Collection("fs.tokens")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return fsFiles, fsTokens, ctx
}

func InitMongoDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Create connect
	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	return client
}
