package data

import (
	"context"
	"errors"
	"time"
	"task_manager/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var taskCollection *mongo.Collection

func Init() {
	// Load MongoDB URI from environment variable
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	uri := os.Getenv("MONGODB_URI")
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	taskCollection = client.Database("test").Collection("tasks") // Collection name updated
	createSampleTasks()
}

func createSampleTasks() {
	// Add sample tasks to the MongoDB collection
	tasks := []interface{}{
		models.Task{
			ID:          "1",
			Title:       "Task 1",
			Description: "Description 1",
			DueDate:     time.Now(),
		},
		models.Task{
			ID:          "2",
			Title:       "Task 2",
			Description: "Description 2",
			DueDate:     time.Now(),
		},
	}

	_, err := taskCollection.InsertMany(context.TODO(), tasks)
	if err != nil {
		log.Fatal(err)
	}
}

func GetTasks() ([]*models.Task, error) {
	var tasks []*models.Task

	cursor, err := taskCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTaskById(id string) (*models.Task, error) {
	var task models.Task
	filter := bson.D{{"id", id}}
	err := taskCollection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return nil, errors.New("task not found")
	}

	return &task, nil
}

func CreateTask(task *models.Task) error {
	_, err := GetTaskById(task.ID)
	if err == nil {
		return errors.New("task already exists")
	}

	_, err = taskCollection.InsertOne(context.TODO(), task)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTask(id string, updatedTask models.Task) error {
	update := bson.M{}

	if updatedTask.Title != "" {
		update["title"] = updatedTask.Title
	}
	if updatedTask.Description != "" {
		update["description"] = updatedTask.Description
	}
	if updatedTask.DueDate != (time.Time{}) {
		update["due_date"] = updatedTask.DueDate
	}

	filter := bson.D{{"id", id}}
	result, err := taskCollection.UpdateOne(context.TODO(), filter, bson.M{"$set": update})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}

func DeleteTask(id string) error {
	// Delete the task by ID
	filter := bson.D{{"id", id}}
	result, err := taskCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}
