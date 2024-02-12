package model

import (
<<<<<<< HEAD
	"server/db"
=======
>>>>>>> 8aebc4dbf2f9d838a919bf0cbcd6025d276794ff
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
<<<<<<< HEAD

func AutoMigrateJawaban() {
	if err := db.DB.AutoMigrate(&Jawaban{}); err != nil {
		panic(err)
	}
}
=======
>>>>>>> 8aebc4dbf2f9d838a919bf0cbcd6025d276794ff
