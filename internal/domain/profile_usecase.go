package domain

type ProfileUsecase interface {
	CreateProfile(profile *Profile) (string, error)
	UpdateProfile(profileID string, profile *Profile, fields []string) (*Profile, error)
	DeleteProfile(profileID string) (*Profile, error)
	GetProfileByID(profileID string) (*Profile, error)
	GetProfileByUserID(userID string) (*Profile, error)
}