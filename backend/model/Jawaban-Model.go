package model

import (
	"time"
)

type Jawaban struct {
	Id_jawaban int       `gorm:"primary_key"`
	Isi        string    `gorm:"type:text" json:"isi"`
	Nilai      uint32    `gorm:"default:0" json:"nilai"`
	CreateAt   time.Time `json:"create_at"`
	File       string    `gorm:"type:varchar(100)" json:"file"`
	SiswaNIS   int
	Siswa      User 
}
