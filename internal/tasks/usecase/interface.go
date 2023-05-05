package usecase

import (
	"backend/internal/tasks/domain"
	"context"
)

type TaskUsecase interface {
	CreateTask(ctx context.Context, req *CreateTaskRequest) (*CreateTaskResponse, error)
	UpdateTask(ctx context.Context, req *UpdateTaskRequest) (*UpdateTaskResponse, error)
	DeleteTask(ctx context.Context, req *DeleteTaskRequest) (*DeleteTaskResponse, error)
	ListAllTask(ctx context.Context, req *ListAllTaskRequest) (*ListAllTaskResponse, error)
	GetTask(ctx context.Context, req *GetTaskRequest) (*GetTaskResponse, error)
}

type CreateTaskRequest struct {
	Task domain.Task
}

type CreateTaskResponse struct{}

type UpdateTaskRequest struct {
	ID     int64
	Task   domain.Task
	UserID string
}

type UpdateTaskResponse struct{}

type DeleteTaskRequest struct {
	ID     int64
	UserID string
}
type DeleteTaskResponse struct{}

type ListAllTaskRequest struct {
	Name   string
	Date   int64
	UserID string
}
type ListAllTaskResponse struct {
	Tasks []*domain.Task
	Total int64
}

type GetTaskRequest struct {
	ID     int64
	UserID string
}

type GetTaskResponse struct {
	Task *domain.Task
}
