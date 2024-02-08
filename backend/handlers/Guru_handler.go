package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Guru struct{
  NIK   int  `json:"nik"`
  Name  string   `json:"name"`
  Passphrase string `json:"passphrase"`
  Email string `json:"email"`
  Gender string `json:"gender"`
  Religion string `json:"religion"`
  Status string `json:"status"`
}

// Create user handler
func CreateGuru(c *gin.Context) {
    var newGuru Guru
  if err := c.ShouldBindJSON(&newGuru); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newGuru.Passphrase), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    // Gunakan password yang di-hash saat memasukkan data ke dalam database
    _, err = dbConn.Exec("INSERT INTO guru (nik, nama, passphrase, email, gender, agama) VALUES ($1, $2, $3, $4, $5, $6)",
        newGuru.NIK, newGuru.Name, hashedPassword, newGuru.Email, newGuru.Gender, newGuru.Religion)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Guru",})

        fmt.Println(err)
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Guru created successfully"})
}
// Update user handler
func UpdateGuru(c *gin.Context) {
    // Mendapatkan ID pengguna dari parameter
    nisStr := c.Param("NIK")

    // Memeriksa apakah nilai "nis" ada
    if nisStr == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "GURU ID is required"})
        return
    }

    // Memeriksa apakah nilai "nis" adalah bilangan bulat
    nis, err := strconv.Atoi(nisStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid GURU ID format"})
        return
    }

    // Mendapatkan data pengguna yang ingin diperbarui dari body permintaan
    var updateguru Guru
    if err := c.ShouldBindJSON(&updateguru); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateguru.Passphrase), bcrypt.DefaultCost)
  if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }


    // Lakukan operasi pembaruan pengguna sesuai dengan kebutuhan aplikasi Anda
    _, err = dbConn.Exec("UPDATE guru SET nama=$1, passphrase=$2, email=$3, gender=$4, agama=$5, status=$6 WHERE nik=$7",
        updateguru.Name, hashedPassword, updateguru.Email, updateguru.Gender, updateguru.Religion, updateguru.Status, nis)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        fmt.Println(err)
        return
    }

    // Respon berhasil
    c.JSON(http.StatusOK, gin.H{"message": "Guru updated successfully"})
}

