package Infrastructure

import (
	"context"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client   *mongo.Client
	database *mongo.Database
)

func InitDB() *mongo.Database {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	database = client.Database("task_manager")
	return database
}

func GetDB() *mongo.Database {
	if database == nil {
		log.Fatal("Database not initialized. Call InitDB first.")
	}
	return database
}

func CollectionNames() (tasksCollection *mongo.Collection, usersCollection *mongo.Collection) {
	db := GetDB()
	tasksCollection = db.Collection("tasks")
	usersCollection = db.Collection("users")
	return
}