package model

import (
	"gorm.io/gorm"
)

type UploadTugas struct {
	gorm.Model

	Name string `gorm:"type:varchar(25)" json:"name"`
	File string `json:"file"`
}
