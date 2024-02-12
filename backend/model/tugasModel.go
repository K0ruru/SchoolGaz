package model

import (
	"server/database"
	"time"
)

type Tugas struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Judul     string    `gorm:"type:varchar(25)" json:"judul"`
	Deskripsi string    `gorm:"type:text" json:"deskripsi"`
	Pembuatan time.Time `json:"pembuatan"`
	Deadline  time.Time `json:"deadline"`
	Status    string    `gorm:"type:varchar(10)" default:"active" json:"status"`
}

func AutoMigrateTugas() {
	if err := database.DB.AutoMigrate(&Tugas{}); err != nil {
		panic(err)
	}
}
