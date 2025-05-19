package postgres

import "github.com/LavaJover/shvark-profile-service/internal/domain"

type DefaultProfileUsecase struct {
	repo domain.ProfileRepository
}

func NewDefaultProfileUsecase(repo domain.ProfileRepository) *DefaultProfileUsecase {
	return &DefaultProfileUsecase{repo}
}

func (uc *DefaultProfileUsecase) CreateProfile(profile *domain.Profile) (string, error) {
	return uc.repo.CreateProfile(profile)
}

func (uc *DefaultProfileUsecase) GetProfileByID(profileID string) (*domain.Profile, error) {
	return uc.repo.GetProfileByID(profileID)
}

func (uc *DefaultProfileUsecase) UpdateProfile(profileID string, profile *domain.Profile, fields []string) (*domain.Profile, error) {
	return uc.repo.UpdateProfile(profileID, profile, fields)
}

func (uc *DefaultProfileUsecase) GetProfileByUserID(userID string) (*domain.Profile, error) {
	return uc.repo.GetProfileByUserID(userID)
} 