package models

import (
	"server/database"
	"time"
)

type Jawaban struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Isi       string    `gorm:"type:text" json:"isi"`
	TugasID   int64     `json:"tugas_id"`
	Tugas     Tugas     `gorm:"foreignKey:TugasID"`
	Pembuatan time.Time `json:"pembuatan"`
	Nilai     *uint32   `gorm:"default:0" json:"nilai"`
}

func AutoMigrateJawaban() {
	if err := database.DB.AutoMigrate(&Jawaban{}); err != nil {
		panic(err)
	}
}
