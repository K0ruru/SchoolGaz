package model
<<<<<<< HEAD:backend/model/tUploadModel.go

import "server/db"
=======
>>>>>>> 8aebc4dbf2f9d838a919bf0cbcd6025d276794ff:backend/model/Tugas-Upload-Model.go

type uploadTugas struct {
	ID   int64   `gorm:"primaryKey" json:"id"`
	Name string  `gorm:"type:varchar(25);default:nil" json:"name"`
	File *string `json:"file" gorm:"default:nil"`
}
<<<<<<< HEAD:backend/model/tUploadModel.go

func AutoMigrateUploadTugas() {
	if err := db.DB.AutoMigrate(&uploadTugas{}); err != nil {
		panic(err)
	}
}
=======
>>>>>>> 8aebc4dbf2f9d838a919bf0cbcd6025d276794ff:backend/model/Tugas-Upload-Model.go
