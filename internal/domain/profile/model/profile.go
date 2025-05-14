package model

import "time"

type Profile struct {
	ID			string
	UserID 		string
	Username 	string
	AvatarURL	string
	Bio			string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

func NewProfile(userID, username, avatarURL, bio string) (*Profile, error){
	return &Profile{
		UserID: userID,
		Username: username,
		AvatarURL: avatarURL,
		Bio: bio,
	}, nil
}