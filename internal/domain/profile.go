package domain

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