package model

type uploadTugas struct {
	ID   int64   `gorm:"primaryKey" json:"id"`
	Name string  `gorm:"type:varchar(25);default:nil" json:"name"`
	File *string `json:"file" gorm:"default:nil"`
}
