package model

type RoleEnum string

const (
	admin RoleEnum = "admin"
	guru  RoleEnum = "guru"
)

type Guru struct {
	NIS            int      `gorm:"primary_key"`
	NamaGuru       string   `gorm:"type:varchar(50)"`
	Email          string   `gorm:"type:varchar(100);uniqueIndex"`
	Passphrase     string   `gorm:"type:varchar(255)"`
	NoTelp         int      `gorm:"type:bigint"`
	Gender         string   `gorm:"type:varchar(25)"`
	Religion       string   `gorm:"type:varchar(25)"`
	ProfilePicture string   `gorm:"type:varchar(50)"`
	Role           RoleEnum `gorm:"type:role_enum;default:'guru'"`
}
