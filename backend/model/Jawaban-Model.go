package model

import (
	"time"
)

type Jawaban struct {
	Id_jawaban int64     `gorm:"primaryKey" json:"id_jawaban"`
	Isi        string    `gorm:"type:text" json:"isi"`
	TugasID    int64     `json:"tugas_id"`
	Tugas      Tugas     `gorm:"foreignKey:TugasID"`
	Pembuatan  time.Time `json:"pembuatan"`
	Nilai      *uint32   `gorm:"default:0" json:"nilai"`
}
