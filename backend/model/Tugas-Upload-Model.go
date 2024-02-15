package model

type UploadTugas struct {
	id_UT int    `gorm:"primary_key"`
	Name  string `gorm:"type:varchar(25)" json:"name"`
	File  string `json:"file"`
}
