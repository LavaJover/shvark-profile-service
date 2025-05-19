package domain

import "time"

type Profile struct {
	ID			string
	UserID 		string
	AvatarURL	string
	TgLink		string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}