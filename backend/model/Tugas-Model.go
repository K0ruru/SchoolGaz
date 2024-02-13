package model

import (
	"time"

	"gorm.io/gorm"
)

type Tugas struct {
	gorm.Model

	Judul     string    `gorm:"type:varchar(25)" json:"judul"`
	Deskripsi string    `gorm:"type:text" json:"deskripsi"`
	Deadline  time.Time `json:"deadline"`
	Id_mapel  *int       
	Kelas     *int
	Mapel     Mapel     `gorm:"foreignkey:Id_mapel"`
	KelasData Kelas 		`gorm:"foreignkey:Kelas"`
}
