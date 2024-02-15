package model

type RoleEnum string

const (
	admin RoleEnum = "admin"
	guru  RoleEnum = "guru"
)

type Guru struct {
	NIS            int      `gorm:"primary_key"`
  NamaGuru       string   `gorm:"type:varchar(50)" json:"nama_guru"`
	Email          string   `gorm:"type:varchar(100);uniqueIndex" json:"email"`
  Passphrase     string   `gorm:"type:varchar(255)" json:"passphrase"`
  NoTelp         int      `gorm:"type:bigint" json:"no_telp"`
	Gender         string   `gorm:"type:varchar(25)" json:"gender"`
	Religion       string   `gorm:"type:varchar(25)" json:"religion"`
	ProfilePicture string   `gorm:"type:varchar(50)" json:"profile_picture"`
	Role           RoleEnum `gorm:"type:role_enum;default:'guru'" json:"role"`
}
