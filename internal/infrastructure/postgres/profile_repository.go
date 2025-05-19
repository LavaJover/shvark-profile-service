package postgres

import (
	"github.com/LavaJover/shvark-profile-service/internal/domain"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type DefaultProfileRepository struct {
	DB	*gorm.DB
}

func NewDefaultProfileRepository(db *gorm.DB) *DefaultProfileRepository {
	return &DefaultProfileRepository{DB: db}
}

func (r *DefaultProfileRepository) CreateProfile(profile *domain.Profile) (string, error) {
	profileModel := &ProfileModel{
		ID: uuid.New().String(),
		UserID: profile.UserID,
		AvatarURL: profile.AvatarURL,
		TgLink: profile.TgLink,
	}
	err := r.DB.Create(profileModel).Error
	if err != nil {
		return "", err
	}
	profile.ID = profileModel.ID
	return profile.ID, nil
}

func (r *DefaultProfileRepository) GetProfileByID(profileID string) (*domain.Profile, error) {
	var profileModel ProfileModel
	if err := r.DB.Where("id = ?", profileID).First(&profileModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrProfileNotFound
		}
		return nil, domain.ErrProfileNotFound
	}

	return &domain.Profile{
		ID: profileModel.ID,
		UserID: profileModel.UserID,
		AvatarURL: profileModel.AvatarURL,
		TgLink: profileModel.TgLink,
	}, nil
}

func (r *DefaultProfileRepository) UpdateProfile(profileID string, profile *domain.Profile, fields []string) (*domain.Profile, error) {
	var profileModel ProfileModel
	if err := r.DB.Where("id = ?", profileID).First(&profileModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrProfileNotFound
		}
		return nil, err
	}

	for _, field := range fields {
		switch field{
		case "AvatarURL":
			profileModel.AvatarURL = profile.AvatarURL
		case "TgLink":
			profileModel.TgLink = profile.TgLink
		}
	}

	if err := r.DB.Save(profileModel).Error; err != nil {
		return nil, err
	}

	return &domain.Profile{
		ID: profileModel.ID,
		UserID: profileModel.UserID,
		AvatarURL: profileModel.AvatarURL,
		TgLink: profileModel.TgLink,
	}, nil
}

func (r *DefaultProfileRepository) DeleteProfile(profileID string) (*domain.Profile, error) {
	var profileModel ProfileModel
	if err := r.DB.Where("id = ?", profileID).First(&profileModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrProfileNotFound
		}

		return nil, err
	}
	if err := r.DB.Delete(profileModel).Error; err != nil {
		return nil, err
	}

	return &domain.Profile{
		ID: profileModel.ID,
		UserID: profileModel.UserID,
		AvatarURL: profileModel.AvatarURL,
		TgLink: profileModel.TgLink,
	}, nil
}

func (r *DefaultProfileRepository) GetProfileByUserID(userID string) (*domain.Profile, error) {
	var profileModel ProfileModel
	if err := r.DB.Where("user_id = ?", userID).First(&profileModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrProfileNotFound
		}
		return nil, err
	}

	return &domain.Profile{
		ID: profileModel.ID,
		UserID: profileModel.UserID,
		AvatarURL: profileModel.AvatarURL,
		TgLink: profileModel.TgLink,
	}, nil
}