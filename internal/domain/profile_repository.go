package domain

type ProfileRepository interface {
	CreateProfile(profile *Profile) (string, error)
	UpdateProfile(profileID string, profile *Profile, fields []string) (*Profile, error)
	DeleteProfile(profileID string) (*Profile, error)
	GetProfileByID(profileID string) (*Profile, error)
}