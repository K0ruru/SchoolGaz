package model

import (
	"gorm.io/gorm"
)

type Jawaban struct {
	gorm.Model

	Isi   string `gorm:"type:text" json:"isi"`
	Nilai uint32 `gorm:"default:0" json:"nilai"`
}
