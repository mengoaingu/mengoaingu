package usecase

import (
	"backend/internal/tasks/repository"
	"context"
)

var _ TaskUsecase = (*taskUsecase)(nil)

type taskUsecase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUsecase(taskrepo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{
		taskRepo: taskrepo,
	}
}

// CreateTask implements TaskUsecase
func (uc *taskUsecase) CreateTask(ctx context.Context, req *CreateTaskRequest) (*CreateTaskResponse, error) {
	err := uc.taskRepo.Save(ctx, &req.Task)
	if err != nil {
		return nil, err
	}
	return &CreateTaskResponse{}, nil
}

// DeleteTask implements TaskUsecase
func (uc *taskUsecase) DeleteTask(ctx context.Context, req *DeleteTaskRequest) (*DeleteTaskResponse, error) {
	err := uc.taskRepo.Delete(ctx, req.ID, req.UserID)
	if err != nil {
		return nil, err
	}
	return &DeleteTaskResponse{}, nil
}

// GetTask implements TaskUsecase
func (uc *taskUsecase) GetTask(ctx context.Context, req *GetTaskRequest) (*GetTaskResponse, error) {
	task, err := uc.taskRepo.FindByID(ctx, req.ID, req.UserID)
	if err != nil {
		return nil, err
	}
	return &GetTaskResponse{
		Task: task,
	}, nil
}

// ListAllTask implements TaskUsecase
func (uc *taskUsecase) ListAllTask(ctx context.Context, req *ListAllTaskRequest) (*ListAllTaskResponse, error) {
	tasks, total, err := uc.taskRepo.Find(ctx, req.Name, req.Date, req.UserID)
	if err != nil {
		return nil, err
	}
	return &ListAllTaskResponse{
		Tasks: tasks,
		Total: total,
	}, nil
}

// UpdateTask implements TaskUsecase
func (uc *taskUsecase) UpdateTask(ctx context.Context, req *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	_, err := uc.taskRepo.Update(ctx, req.ID, req.UserID, &req.Task)
	if err != nil {
		return nil, err
	}
	return &UpdateTaskResponse{}, nil
}
