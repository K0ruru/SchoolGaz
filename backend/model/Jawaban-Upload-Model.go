package model

type UploadJawaban struct {
	id_UJ int    `gorm:"primary_key"`
	Name  string `gorm:"type:varchar(25)" json:"name"`
	File  string `json:"file"`
}
