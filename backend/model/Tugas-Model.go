package model

import (
	"time"
)

type Tugas struct {
	Id_tugas  int64     `gorm:"primaryKey" json:"id_tugas"`
	Judul     string    `gorm:"type:varchar(25)" json:"judul"`
	Deskripsi string    `gorm:"type:text" json:"deskripsi"`
	Pembuatan time.Time `json:"pembuatan"`
	Deadline  time.Time `json:"deadline"`
	Status    string    `gorm:"type:varchar(10)" default:"active" json:"status"`
}
