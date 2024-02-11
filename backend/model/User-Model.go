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
	NoTelp         int64      `gorm:"type:bigint" json:"No_telp"`
	Gender         string     `gorm:"type:varchar(25)"`
	Religion       string     `gorm:"type:varchar(25)"`
	ProfilePicture string     `gorm:"type:varchar(50)"`
	Status         StatusEnum `gorm:"type:status_enum;default:'pending'"`
	KelasID        uint     
	Kelas          Kelas      `gorm:"foreignkey:KelasID"`
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

    // Check if foreign key constraint already exists for Kelas model
    var count int64
    if err := db.Raw("SELECT count(*) FROM information_schema.table_constraints WHERE constraint_type = 'FOREIGN KEY' AND table_name = 'kelas' AND constraint_name = 'kelas_walas_id_gurus_nis_foreign'").Row().Scan(&count); err != nil {
        return err
    }

    // Add foreign key constraint for Kelas model if it doesn't exist
    if count == 0 {
        if err := db.Exec("ALTER TABLE kelas ADD CONSTRAINT kelas_walas_id_gurus_nis_foreign FOREIGN KEY (walas_id) REFERENCES gurus(NIS) ON DELETE CASCADE ON UPDATE CASCADE").Error; err != nil {
            return err
        }
    }

    // Check if foreign key constraint already exists for User model
    if err := db.Raw("SELECT count(*) FROM information_schema.table_constraints WHERE constraint_type = 'FOREIGN KEY' AND table_name = 'user' AND constraint_name = 'user_kelas_id_foreign'").Row().Scan(&count); err != nil {
        return err
    }

    // Add foreign key constraint for User model if it doesn't exist
    if count == 0 {
        if err := db.Exec("ALTER TABLE users ADD CONSTRAINT user_kelas_id_foreign FOREIGN KEY (kelas_id) REFERENCES kelas(id_kelas) ON DELETE CASCADE ON UPDATE CASCADE").Error; err != nil {
            return err
        }
    }

    return nil
}


