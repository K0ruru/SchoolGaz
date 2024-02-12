package model

type uploadTugas struct {
	Id_ut int64  `gorm:"primaryKey" json:"id_ut"`
	Name  string `gorm:"type:varchar(25)" json:"name"`
	File  string `json:"file"`
}
