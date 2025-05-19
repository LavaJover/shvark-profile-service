package postgres

import "github.com/LavaJover/shvark-profile-service/internal/domain"

type DefaultProfileUsecase struct {
	Repo domain.ProfileRepository
}

func NewDefaultProfileUsecase(repo domain.ProfileRepository) *DefaultProfileUsecase {
	return &DefaultProfileUsecase{repo}
}

func (uc *DefaultProfileUsecase) CreateProfile(profile *domain.Profile) (string, error) {
	return uc.Repo.CreateProfile(profile)
}

func (uc *DefaultProfileUsecase) GetProfileByID(profileID string) (*domain.Profile, error) {
	return uc.Repo.GetProfileByID(profileID)
}

func (uc *DefaultProfileUsecase) UpdateProfile(profileID string, profile *domain.Profile, fields []string) (*domain.Profile, error) {
	return uc.Repo.UpdateProfile(profileID, profile, fields)
}

func (uc *DefaultProfileUsecase) GetProfileByUserID(userID string) (*domain.Profile, error) {
	return uc.Repo.GetProfileByUserID(userID)
} 

func (uc *DefaultProfileUsecase) DeleteProfile(profileID string) (*domain.Profile, error) {
	return uc.Repo.DeleteProfile(profileID)
}