package model

import (
	"github.com/jinzhu/gorm"
)

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
}

func AutoMigrate(db *gorm.DB) error {
	// Check if the custom enum type exists
	var typeExists bool
	if err := db.Raw("SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_enum')").Row().Scan(&typeExists); err != nil {
		return err
	}

	// If the type does not exist, create it
	if !typeExists {
		if err := db.Exec("CREATE TYPE status_enum AS ENUM ('pending', 'active')").Error; err != nil {
			return err
		}
	}

	// AutoMigrate other models
	if err := db.AutoMigrate(&Guru{}).Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	if err := db.AutoMigrate(&Kelas{}).Error; err != nil {
		return err
	}

	// Add index for Guru model
	db.Model(&Guru{}).AddIndex("idx_gurus_role", "role")

	return nil
}

