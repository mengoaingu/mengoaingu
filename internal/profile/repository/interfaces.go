package repository

import (
	"backend/internal/profile/domain"
	"context"
)

type ProfileRepository interface {
	Find(ctx context.Context, profile_id int64) (*domain.User, error)
	FindByUserID(ctx context.Context, user_id string) (*domain.User, error)
	Save(ctx context.Context, profile *domain.User) error
}
