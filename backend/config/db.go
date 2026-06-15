package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client
var UserCollection *mongo.Collection
var TaskCollection *mongo.Collection
var TeamCollection *mongo.Collection

func ConnectDB() {
	_ = godotenv.Load()
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	DB = client
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "taskforge"
	}
	UserCollection = client.Database(dbName).Collection("users")
	TaskCollection = client.Database(dbName).Collection("tasks")
	TeamCollection = client.Database(dbName).Collection("teams")
}
