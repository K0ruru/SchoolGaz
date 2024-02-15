package handlers

import (
	"fmt"
	"net/http"
	"server/model"
	"strconv"

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

func GetKelas(c *gin.Context) {
    var kelas model.Kelas
    ID := c.Param("id_kelas")
    
    kelasID, err := strconv.Atoi(ID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id_kelas"})
        fmt.Println(err)
        return
    }
    
    // Fetch the kelas with the given ID
    if err := dbConn.First(&kelas, kelasID).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }
    
    // Fetch associated Walas for the Kelas
    var walas model.Guru
    if err := dbConn.First(&walas, kelas.WalasNIS).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }
    // Exclude passphrase from Walas object
    walas.Passphrase = ""
    kelas.Walas = walas
    
    c.JSON(http.StatusOK, kelas)
}

func CreateKelas(c *gin.Context) {
    var NewKelas model.Kelas

    // Bind JSON body to Kelas struct
    if err := c.ShouldBindJSON(&NewKelas); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        fmt.Println(err  )
        return
    }

    // Create new record for Kelas in the database
    if err := dbConn.Create(&NewKelas).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    // Respond with the newly created Kelas
    c.JSON(http.StatusOK, NewKelas)
}

func UpdateKelas(c *gin.Context) {
    if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
		return
	}

    kelasID := c.Param("id_kelas")

    _, err := strconv.Atoi(kelasID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kelasID"})
        return
    }
    var updateKelas model.Kelas

    if err := dbConn.Where("Id_kelas = $1", kelasID).First(&updateKelas).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
    }

    if err := c.ShouldBindJSON(&updateKelas); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }

    if err := dbConn.Save(&updateKelas).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
    }

    c.JSON(http.StatusOK, updateKelas)
}

func DeleteKelas(c *gin.Context) {
    if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
		return
	}

    kelasID := c.Param("id_kelas")

    _, err := strconv.Atoi(kelasID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kelasID"})
        return
    }
    var deleteKelas model.Kelas

    if err := dbConn.Where("Id_kelas = $1", kelasID).Delete(&deleteKelas).Error; err!= nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
    }

    c.JSON(http.StatusOK, gin.H{"success": "kelas deleted"})
}
