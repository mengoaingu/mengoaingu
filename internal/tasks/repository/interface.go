package repository

import (
	"backend/internal/tasks/domain"
	"context"
)

type TaskRepository interface {
	Save(ctx context.Context, task *domain.Task) error
	Find(ctx context.Context, name string, date int64, user_id string) ([]*domain.Task, int64, error)
	FindByID(ctx context.Context, task_id int64, user_id string) (*domain.Task, error)
	Update(ctx context.Context, task_id int64, user_id string, task *domain.Task) (*domain.Task, error)
	Delete(ctx context.Context, task_id int64, user_id string) error
}
