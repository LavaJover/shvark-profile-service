package domain

import "google.golang.org/protobuf/types/known/fieldmaskpb"

type ProfileUsecase interface {
	CreateProfile(profile *Profile) (string, error)
	UpdateProfile(profileID string, profile *Profile, updateMask *fieldmaskpb.FieldMask) (*Profile, error)
	DeleteProfile(profileID string) (*Profile, error)
	GetProfileByID(profileID string) (*Profile, error)
	GetProfileByUserID(userID string) (*Profile, error)
}