package handlers

import (
	"fmt"
	"net/http"
	"strconv"

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
// Update user handler
func Updatekelas(c *gin.Context) {
    // Mendapatkan ID pengguna dari parameter
    id_kelas := c.Param("id_kelas")

    // Memeriksa apakah nilai "nis" ada
    if id_kelas == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "kelas ID is required"})
        return
    }

    // Memeriksa apakah nilai "nis" adalah bilangan bulat
    ID, err := strconv.Atoi(id_kelas)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kelas ID format"})
        return
    }

    // Mendapatkan data pengguna yang ingin diperbarui dari body permintaan
    var UpdateKelas Kelas
    if err := c.ShouldBindJSON(&UpdateKelas); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        fmt.Println(err)
        return
    }

    // Lakukan operasi pembaruan pengguna sesuai dengan kebutuhan aplikasi Anda
    _, err = dbConn.Exec("UPDATE kelas SET siswa=$1, walas=$2, nama_kelas=$3  WHERE id_kelas=$4",
    UpdateKelas.Siswa,UpdateKelas.Walas,UpdateKelas.Nama_kelas, ID) 
  
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update kelas"})
        fmt.Println(err)
        return
    }

    // Respon berhasil
    c.JSON(http.StatusOK, gin.H{"message": "kelas updated successfully"})
}
// Get all users handler
func Getkelas(c *gin.Context) {
    rows, err := dbConn.Query("SELECT id_kelas, siswa, walas, nama_kelas FROM kelas")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch kelas"})
        return
    }
    defer rows.Close()

    var kelass []Kelas
    for rows.Next() {
        var kelas Kelas
        if err := rows.Scan(&kelas.ID, &kelas.Siswa,&kelas.Walas, &kelas.Nama_kelas); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan users"})
            fmt.Println(err)
            return
        }
        kelass = append(kelass, kelas)
    }

    if err := rows.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan fetch users"})
        return
    }

    c.JSON(http.StatusOK, kelass)
}
// DeleteUser handles the deletion of a user
func DeleteKelas(c *gin.Context) {
    ID := c.Param("id_kelas")
    
    // Memeriksa apakah nilai "nis" ada
    if ID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "kelas ID is required"})
        return
    }

    // Memeriksa apakah nilai "nis" adalah bilangan bulat
    id_kelas, err := strconv.Atoi(ID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
        return
    }
   _, err = dbConn.Exec("DELETE FROM kelas WHERE id_kelas=$1", id_kelas)
      if err != nil {
         fmt.Println("Failed to delete kelas: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete kelas"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "kelas deleted successfully"})
}
