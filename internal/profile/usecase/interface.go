package usecase

import (
	"backend/internal/profile/domain"
	"context"
)

type ProfileUsecase interface {
	GetProfile(context.Context, *GetProfileRequest) (*GetProfileResponse, error)
	CreateProfile(context.Context, *CreateProfileRequest) (*CreateProfileResponse, error)
}

type GetProfileRequest struct {
	UserID string
}

type GetProfileResponse struct {
	Profile *domain.User
}

type CreateProfileRequest struct {
	Profile *domain.User
}

type CreateProfileResponse struct {
	Profile *domain.User
}
