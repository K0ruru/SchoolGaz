// models/tugas.go
package models

import (
    "time"
    "server/database"
)

type Tugas struct {
    ID        int64     `gorm:"primaryKey" json:"id"`
    Judul     string    `gorm:"type:varchar(25)" json:"judul"`
    Deskripsi string    `gorm:"type:text" json:"deskripsi"`
    Pembuatan time.Time `json:"pembuatan"`
    Deadline  time.Time `json:"deadline"`
    Status    string    `gorm:"type:varchar(10);check:status IN ('active', 'notActive')" json:"status"`
}

// AutoMigrateTugas automatically migrates the Tugas table
func AutoMigrateTugas() {
    if err := database.DB.AutoMigrate(&Tugas{}); err != nil {
        panic(err)
    }
}
