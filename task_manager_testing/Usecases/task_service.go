package Usecases

import (
	"context"
	"errors"
	"task_manager_clean_architecture/Domain"
	"task_manager_clean_architecture/Repositories"
	
)


type TaskUsecase struct {
	taskRepo Repositories.TaskRepository
}

func NewTaskUsecase(taskRepo Repositories.TaskRepository) *TaskUsecase {
	return &TaskUsecase{
		taskRepo: taskRepo,
	}
}

func (uc *TaskUsecase) GetTasks() ([]*Domain.Task, error) {
	return uc.taskRepo.GetTasks(context.TODO())
}

func (uc *TaskUsecase) GetTaskByID(id string) (*Domain.Task, error) {
	return uc.taskRepo.GetTaskByID(context.TODO(), id)
}

func (uc *TaskUsecase) CreateTask(task *Domain.Task) error {
	existingTask, _ := uc.taskRepo.GetTaskByID(context.TODO(), task.ID)
	if existingTask != nil {
		return errors.New("task already exists")
	}
	return uc.taskRepo.CreateTask(context.TODO(), task)
}

func (uc *TaskUsecase) UpdateTask(id string, updatedTask Domain.Task) error {
	_, err := uc.taskRepo.GetTaskByID(context.TODO(), id)
	if err != nil {
		return errors.New("task not found")
	}
	return uc.taskRepo.UpdateTask(context.TODO(), id, updatedTask)
}

func (uc *TaskUsecase) DeleteTask(id string) error {
	_, err := uc.taskRepo.GetTaskByID(context.TODO(), id)
	if err != nil {
		return errors.New("task not found")
	}
	return uc.taskRepo.DeleteTask(context.TODO(), id)
}



