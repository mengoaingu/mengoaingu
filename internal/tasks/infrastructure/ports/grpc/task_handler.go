package grpc

import (
	"backend/internal/tasks/domain"
	"backend/internal/tasks/usecase"
	"backend/pkg"
	"backend/pkg/middleware"
	"backend/proto/gen"
	"context"
	"errors"
)

type TaskGRPCServer struct {
	taskUC     usecase.TaskUsecase
	authClient *middleware.Authorization
}

func NewTaskGRPCServer(app *pkg.App, taskUC usecase.TaskUsecase, authClient *middleware.Authorization) *TaskGRPCServer {
	taskGrpcServer := &TaskGRPCServer{
		taskUC:     taskUC,
		authClient: authClient,
	}
	gen.RegisterTaskServiceServer(app.GrpcServer, taskGrpcServer)

	return taskGrpcServer
}

var _ gen.TaskServiceServer = (*TaskGRPCServer)(nil)

// CreateTask implements gen.TaskServiceServer
func (h *TaskGRPCServer) CreateTask(ctx context.Context, request *gen.CreateTaskRequest) (*gen.CreateTaskResponse, error) {
	token, err := h.authClient.AuthorizeN(ctx)
	if err != nil {
		return nil, errors.New("you don't have permission")
	}

	task := request.Task

	_, err = h.taskUC.CreateTask(ctx, &usecase.CreateTaskRequest{
		Task: domain.Task{
			Name:        task.Name,
			Description: task.Description,
			Status:      domain.TaskStatus(task.Status),
			Priority:    domain.Priority(task.Priority),
			UserID:      token.UID,
		},
	})

	if err != nil {
		return nil, errors.New("Failed to create task: " + err.Error())
	}

	return &gen.CreateTaskResponse{}, nil
}

// DeleteTask implements gen.TaskServiceServer
func (h *TaskGRPCServer) DeleteTask(ctx context.Context, request *gen.DeleteTaskRequest) (*gen.DeleteTaskResponse, error) {
	token, err := h.authClient.AuthorizeN(ctx)
	if err != nil {
		return nil, errors.New("you don't have permission")
	}

	taskId := request.TaskId
	_, err = h.taskUC.DeleteTask(ctx, &usecase.DeleteTaskRequest{
		ID:     taskId,
		UserID: token.UID,
	})
	if err != nil {
		return nil, errors.New("Failed to delete task: " + err.Error())
	}
	return &gen.DeleteTaskResponse{}, nil
}

// GetListTask implements gen.TaskServiceServer
func (h *TaskGRPCServer) GetListTask(ctx context.Context, request *gen.GetListTaskRequest) (*gen.GetListTaskResponse, error) {
	token, err := h.authClient.AuthorizeN(ctx)
	if err != nil {
		return nil, errors.New("you don't have permission")
	}

	resp, err := h.taskUC.ListAllTask(ctx, &usecase.ListAllTaskRequest{
		Name:   request.Name,
		Date:   request.Date,
		UserID: token.UID,
	})
	if err != nil {
		return nil, errors.New("Failed to get task: " + err.Error())
	}
	var pbListTask []*gen.Task
	for _, t := range resp.Tasks {
		pbListTask = append(pbListTask, &gen.Task{
			Id:          int64(t.ID),
			Name:        t.Name,
			Description: t.Description,
			Priority:    gen.Priority(t.Priority),
			Status:      gen.TaskStatus(t.Status),
		})
	}
	return &gen.GetListTaskResponse{
		Data: &gen.GetListTaskResponse_Data{
			Tasks: pbListTask,
			Total: resp.Total,
		},
	}, nil
}

// GetTask implements gen.TaskServiceServer
func (h *TaskGRPCServer) GetTask(ctx context.Context, request *gen.GetTaskRequest) (*gen.GetTaskResponse, error) {
	token, err := h.authClient.AuthorizeN(ctx)
	if err != nil {
		return nil, errors.New("you don't have permission")
	}

	taskId := request.TaskId
	resp, err := h.taskUC.GetTask(ctx, &usecase.GetTaskRequest{
		ID:     taskId,
		UserID: token.UID,
	})
	if err != nil {
		return nil, errors.New("Failed to get task: " + err.Error())
	}

	return &gen.GetTaskResponse{
		Data: &gen.GetTaskResponse_Data{
			Task: &gen.Task{
				Id:          int64(resp.Task.ID),
				Name:        resp.Task.Name,
				Description: resp.Task.Description,
				Priority:    gen.Priority(resp.Task.Priority),
				Status:      gen.TaskStatus(resp.Task.Status),
			},
		},
	}, nil
}

// UpdateTask implements gen.TaskServiceServer
func (h *TaskGRPCServer) UpdateTask(ctx context.Context, request *gen.UpdateTaskRequest) (*gen.UpdateTaskResponse, error) {
	token, err := h.authClient.AuthorizeN(ctx)
	if err != nil {
		return nil, errors.New("you don't have permission")
	}

	task := request.Task
	_, err = h.taskUC.UpdateTask(ctx, &usecase.UpdateTaskRequest{
		ID: request.TaskId,
		Task: domain.Task{
			Name:        task.Name,
			Description: task.Description,
			Status:      domain.TaskStatus(task.Status),
			Priority:    domain.Priority(task.Priority),
		},
		UserID: token.UID,
	})
	if err != nil {
		return nil, errors.New("Failed to update task: " + err.Error())
	}
	return &gen.UpdateTaskResponse{}, nil
}
