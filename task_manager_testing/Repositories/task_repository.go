package Repositories

import (
	"context"
	"task_manager_clean_architecture/Domain"
)

type TaskRepository interface {
	GetTasks(ctx context.Context) ([]*Domain.Task, error)
	GetTaskByID(ctx context.Context, id string) (*Domain.Task, error)
	CreateTask(ctx context.Context, task *Domain.Task) error
	UpdateTask(ctx context.Context, id string, updatedTask Domain.Task) error
	DeleteTask(ctx context.Context, id string) error
}