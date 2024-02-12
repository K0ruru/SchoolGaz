package model

type StatusEnum string

const (
	Pending StatusEnum = "pending"
	Active  StatusEnum = "active"
)

type User struct {
	NIS            int        `gorm:"primary_key"`
	Nama           string     `gorm:"type:varchar(50)"`
	Email          string     `gorm:"type:varchar(100);unique_index"`
	Passphrase     string     `gorm:"type:varchar(255)"`
	No_telp        int64      `gorm:"type:bigint" json:"No_telp"`
	Gender         string     `gorm:"type:varchar(25)"`
	Religion       string     `gorm:"type:varchar(25)"`
	Profilepicture string     `gorm:"type:varchar(50)"`
	Status         StatusEnum `gorm:"type:status_enum;default:'pending'"`
	Kelas          *int
}
