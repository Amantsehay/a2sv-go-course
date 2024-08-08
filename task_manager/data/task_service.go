package data

import (
	"errors"
	"time"
	"task_manager/models"
)

var tasks = []*models.Task{}


func Init(){
	tasks = append(tasks, &models.Task{
		ID: "1",
		Title: "Task 1",
		Description: "Description 1",
		DueDate: time.Now(),
	})
	tasks = append(tasks, &models.Task{
		ID: "2",
		Title: "Task 2",
		Description: "Description 2",
		DueDate: time.Now(),
	})
}



func GetTasks() []*models.Task{
	return tasks
}
func GetTasksById(id string) (*models.Task, error){
	for _, t := range tasks{
		if t.ID == id{
			return t, nil
		}
	}
	return nil,  errors.New("task not found")
}

func CreateTask(task *models.Task) error{
	existingTask, _ := GetTasksById(task.ID)
	if existingTask != nil{
		return errors.New("task already exists")
	}
	tasks = append(tasks, task)
	return nil
}

func UpdateTask(id string, task models.Task) error{
	var updatedTask *models.Task
	for i, task := range tasks{
		if task.ID == id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if updatedTask.DueDate != (time.Time{}) {
				tasks[i].DueDate = updatedTask.DueDate
			}
	
			return nil
		}
	}
	return errors.New("task not found")
}

func DeleteTask(id string) error{
	for i, t := range tasks{
		if t.ID == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}

