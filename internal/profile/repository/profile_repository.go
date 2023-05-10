package repository

import (
	"backend/internal/profile/domain"
	"context"

	"gorm.io/gorm"
)

type ProfileDB *gorm.DB

var _ ProfileRepository = (*profileRepository)(nil)

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(DB ProfileDB) ProfileRepository {
	return &profileRepository{
		db: DB,
	}
}

// Find implements ProfileRepository
func (r *profileRepository) Find(ctx context.Context, profile_id int64) (*domain.User, error) {
	var task domain.User
	err := r.db.WithContext(ctx).Where("id = ?", profile_id).Find(&task).Error
	return &task, err
}

// FindByUserID implements ProfileRepository
func (r *profileRepository) FindByUserID(ctx context.Context, user_id string) (*domain.User, error) {
	var profile domain.User
	err := r.db.WithContext(ctx).Where("uid = ?", user_id).Find(&profile).Error
	return &profile, err
}

// Save implements ProfileRepository
func (r *profileRepository) Save(ctx context.Context, profile *domain.User) error {
	err := r.db.WithContext(ctx).Create(&profile).Error
	return err
}
