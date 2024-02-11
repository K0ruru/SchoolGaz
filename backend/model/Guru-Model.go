package model



type Guru struct {
    NIS            int    `gorm:"primary_key"`
    NamaGuru       string `gorm:"type:varchar(50)"`
    Email          string `gorm:"type:varchar(100);uniqueIndex"`
    Passphrase     string `gorm:"type:varchar(255)"`
    NoTelp         int    `gorm:"type:bigint"`
    Gender         string `gorm:"type:varchar(25)"`
    Religion       string `gorm:"type:varchar(25)"`
    ProfilePicture string `gorm:"type:varchar(50)"`
  Role           Level  `gorm:"type:role"`
}



type Level string

const(
  guru Level = "guru"
  admin Level = "admin"

)

