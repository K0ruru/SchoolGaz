package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Kelas struct{
  ID          int       `json:"id_kelas"`
  Walas       int       `json:"walas"`
  Nama_kelas  string    `json:"nama_kelas"`
  NamaWalas   string    `json:"nama_walas"` //Gw buat filed buat nyimpan nama walasnya
}

func GetAllKelas(c *gin.Context) {
    rows, err := dbConn.Query("SELECT kelas.*, users.nama FROM kelas JOIN users ON kelas.walas = users.nis") //Ini kurleb querry nya bilang hey ambil nama dari tabel user 
                                                                                                            //terus join usernya dimana data kolom walas sama dengan kolom nis di tabel users

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan data kelas"})
        fmt.Println(err)
        return
    }
    defer rows.Close()

    var kelasKelas []Kelas

    for rows.Next() {
        var kelas Kelas                                                //dimasukin ke field kelas tadi
        if err:= rows.Scan(&kelas.ID, &kelas.Walas, &kelas.Nama_kelas, &kelas.NamaWalas); err != nil {
             c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ngescan"})
             fmt.Println(err)
        }
        kelasKelas = append(kelasKelas, kelas)
    }

    if err := rows.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error" : "Gagal mendapatkan user yang di scan"})
        fmt.Println(err)
    }

    c.JSON(http.StatusOK, kelasKelas)
}