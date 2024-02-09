package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Kelas struct{
  ID          int       `json:"id_kelas"`
  Siswa       int    `json:"siswa"`
  Walas       int        `json:"walas"`
  Nama_kelas  string    `json:"nama_kelas"`
}

func Createkelas(c * gin.Context){
  var newKelas Kelas
  if err := c.ShouldBindJSON(&newKelas); err != nil{
    c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
    fmt.Println(err)
  } 
  _, err := dbConn.Exec("INSERT INTO kelas (siswa, walas, nama_kelas) VALUES ($1,$2,$3)",
     newKelas.Siswa, newKelas.Walas, newKelas.Nama_kelas)
    if err != nil{
        c.JSON(http.StatusBadRequest, gin.H{"error": " salah insert keknya bang"})
        fmt.Println(err)
        return
    }
  c.JSON(http.StatusCreated, gin.H{"error": "Kelas Created Succes"})
}
