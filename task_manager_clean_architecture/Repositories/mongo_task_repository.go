package Repositories

import (
	"context"
	"errors"
	"task_manager_clean_architecture/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"task_manager_clean_architecture/Infrastructure"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTaskRepository struct {
	collection *mongo.Collection
}

func NewMongoTaskRepository() *MongoTaskRepository {
	db := Infrastructure.GetDB()
	tasksCollection := db.Collection("tasks")
	return &MongoTaskRepository{collection: tasksCollection}
}

func (r *MongoTaskRepository) GetTasks(ctx context.Context) ([]*Domain.Task, error) {
	var tasks []*Domain.Task
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var task Domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *MongoTaskRepository) GetTaskByID(ctx context.Context, id string) (*Domain.Task, error) {
	var task Domain.Task
	filter := bson.D{{"id", id}}
	err := r.collection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		return nil, errors.New("task not found")
	}
	return &task, nil
}

func (r *MongoTaskRepository) CreateTask(ctx context.Context, task *Domain.Task) error {
	_, err := r.GetTaskByID(ctx, task.ID)
	if err == nil {
		return errors.New("task already exists")
	}
	_, err = r.collection.InsertOne(ctx, task)
	return err
}

func (r *MongoTaskRepository) UpdateTask(ctx context.Context, id string, updatedTask Domain.Task) error {
	update := bson.M{}
	if updatedTask.Title != "" {
		update["title"] = updatedTask.Title
	}
	if updatedTask.Description != "" {
		update["description"] = updatedTask.Description
	}
	if !updatedTask.DueDate.IsZero() {
		update["due_date"] = updatedTask.DueDate
	}

	filter := bson.D{{"id", id}}
	result, err := r.collection.UpdateOne(ctx, filter, bson.M{"$set": update})
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (r *MongoTaskRepository) DeleteTask(ctx context.Context, id string) error {
	filter := bson.D{{"id", id}}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
