package postgres

type ProfileModel struct {
	ID			string	`gorm:"primaryKey;type:uuid"`
	UserID		string	`gorm:"type:uuid;unique;not null"`
	AvatarURL	string
	TgLink		string
}