package handlers

import (
	"fmt"
	"net/http"
	"server/model"

	"github.com/gin-gonic/gin"
)





func GetAllKelas(c *gin.Context) {
    if dbConn == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "error connection or from the server error"})
        return
    }

    var AllKelas []model.Kelas
    if err := dbConn.Preload("Walas").Find(&AllKelas).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find kelas"})
        return
    }

    c.JSON(http.StatusOK, AllKelas)
}

func CreateKelas(c *gin.Context) {
    var NewKelas model.Kelas

    // Bind JSON body to Kelas struct
    if err := c.ShouldBindJSON(&NewKelas); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Create new record in the database
    if err := dbConn.Create(&NewKelas).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    // Respond with the newly created Kelas
    c.JSON(http.StatusOK, NewKelas)
}
/*
func GetKelas(c *gin.Context) {
    id := c.Param("id") // ambil id nya dari url 

    //cek kalau id nya ada
    kelasID, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_kelas"})
        return
    }

    row := dbConn.QueryRow("SELECT kelas.*, users.nama FROM kelas JOIN users ON kelas.walas = users.nis WHERE kelas.id_kelas = $1", kelasID)

    var kelas Kelas

    err = row.Scan(&kelas.ID, &kelas.Walas, &kelas.Nama_kelas, &kelas.NamaWalas)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Kelas not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get kelas"})
        fmt.Println(err)
        return
    }

    c.JSON(http.StatusOK, kelas)
}

func CreateKelas(c *gin.Context) {
    var kelasBaru Kelas

    if err := c.ShouldBindJSON(&kelasBaru); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        fmt.Println(err)
    }

    _, err := dbConn.Exec("INSERT INTO kelas(walas, nama_kelas) VALUES ($1, $2)", kelasBaru.Walas, kelasBaru.Nama_kelas)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    c.JSON(http.StatusCreated, kelasBaru)
}

func EditKelas(c *gin.Context) {
    id := c.Param("id")

    //cek id nya
    kelasID, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_kelas"})
        return
    }

    // ambil data kelasnya yang lama
    existingKelas := Kelas{}
    err = dbConn.QueryRow("SELECT id_kelas, walas, nama_kelas FROM kelas WHERE id_kelas = $1", kelasID).
        Scan(&existingKelas.ID, &existingKelas.Walas, &existingKelas.Nama_kelas)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Kelas not found"})
            return
        }
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get kelas"})
    fmt.Println(err)
    return
}

    // Bind data jsonnya ke struct
    if err := c.ShouldBindJSON(&existingKelas); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    // execute updatenya dengan data yang ka ganti aman aman aja
    _, err = dbConn.Exec("UPDATE kelas SET walas=$1, nama_kelas=$2 WHERE id_kelas=$3", existingKelas.Walas, existingKelas.Nama_kelas, kelasID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    existingKelas.ID = kelasID
    c.JSON(http.StatusOK, existingKelas)
}

func DeleteKelas(c *gin.Context) {
    id := c.Param("id")


    kelasID, err := strconv.Atoi(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_kelas"})
        return
    }


    _, err = dbConn.Exec("DELETE FROM kelas WHERE id_kelas = $1", kelasID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete kelas"})
        fmt.Println(err)
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Kelas deleted successfully"})
}

*/
