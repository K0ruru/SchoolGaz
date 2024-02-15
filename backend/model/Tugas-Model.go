package model

import (
	"time"
)

type Tugas struct {
	Id_tugas   int       `gorm:"primary_key"`
	Judul      string    `gorm:"type:varchar(25)" json:"judul"`
	Deskripsi  string    `gorm:"type:text" json:"deskripsi"`
	Deadline   time.Time `json:"deadline"`
	CreateAt   time.Time `json:"create_at"`
	File       string    `gorm:"type:varchar(100)" json:"file"`
	Id_mapel   int
	Mapel      Mapel
	Id_jawaban int
	Jawaban    Jawaban
	WalasNIS   int
	Walas      Guru
}
