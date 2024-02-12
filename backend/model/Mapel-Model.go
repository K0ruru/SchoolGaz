package model

type Mapel struct {
  Id_mapel int `gorm:"primary_key"`
  Nama_mapel string `gorm:"type:varchar(50)"`
}
