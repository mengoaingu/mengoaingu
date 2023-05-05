package usecase

import (
	"backend/internal/profile/repository"
	"context"
)

var _ ProfileUsecase = (*profileUsecase)(nil)

type profileUsecase struct {
	profileRepo repository.ProfileRepository
}

func NewProfileUsecase(profileRepo repository.ProfileRepository) ProfileUsecase {
	return &profileUsecase{
		profileRepo: profileRepo,
	}
}

// CreateProfile implements ProfileUsecase
func (uc *profileUsecase) CreateProfile(ctx context.Context, req *CreateProfileRequest) (*CreateProfileResponse, error) {
	err := uc.profileRepo.Save(ctx, req.Profile)
	if err != nil {
		return nil, err
	}
	return &CreateProfileResponse{
		Profile: req.Profile,
	}, nil
}

// GetProfile implements ProfileUsecase
func (uc *profileUsecase) GetProfile(ctx context.Context, req *GetProfileRequest) (*GetProfileResponse, error) {
	profile, err := uc.profileRepo.FindByUserID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	return &GetProfileResponse{
		Profile: profile,
	}, nil
}
