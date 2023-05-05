package repository

import (
	"backend/internal/profile/domain"
	"context"
)

type ProfileRepository interface {
	Find(ctx context.Context, profile_id int64) (*domain.Profile, error)
	FindByUserID(ctx context.Context, user_id string) (*domain.Profile, error)
	Save(ctx context.Context, profile *domain.Profile) error
}
