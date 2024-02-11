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
    Kelas          kelas  `gorm:"foreignkey:KelasID"` // Define the association
}

type kelas struct {
    IDKelas int // Assuming this is the primary key of Kelas
    // Other fields in Kelas struct
}

func AutoMigrate(db *gorm.DB) error {
    // Automigrate Guru
    if err := db.AutoMigrate(&Guru{}).Error; err != nil {
        return err
    }

    // Automigrate User
    if err := db.AutoMigrate(&User{}).Error; err != nil {
        return err
    }

    // Automigrate Kelas
    if err := db.AutoMigrate(&Kelas{}).Error; err != nil {
        return err
    }

    // Add foreign key constraint for User
    if err := db.Model(&User{}).AddForeignKey("kelas_id", "kelas(id_kelas)", "RESTRICT", "RESTRICT").Error; err != nil {
        return err
    }

    return nil
}

