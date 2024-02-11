package model

import (
	

	"github.com/jinzhu/gorm"
)

type User struct {
    NIS            int    `gorm:"primary_key"`
    Nama           string `gorm:"type:varchar(50)"`
    Email          string `gorm:"type:varchar(100);unique_index"`
    Passphrase     string `gorm:"type:varchar(255)"`
    No_telp        int64  `gorm:"type:bigint" json:"No_telp"`
    Gender         string `gorm:"type:varchar(25)"`
    Religion       string `gorm:"type:varchar(25)"`
    Profilepicture string `gorm:"type:varchar(50)"`
}

func AutoMigrate(db *gorm.DB) error {
    if err := db.AutoMigrate(&Guru{}).Error; err != nil {
        return err
    }
    if err := db.AutoMigrate(&User{}).Error; err != nil {
        return err
    }
    if err := db.AutoMigrate(&Kelas{}).Error; err != nil {
        return err // Return the error immediately
    }
  db.Model(&Guru{}).AddIndex("idx_gurus_role", "role")
    return nil
}

