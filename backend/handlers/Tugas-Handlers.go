package handlers

import (
	"fmt"
	"net/http"
	"server/db"
	"server/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func IndexTugas(c *gin.Context) {
	var tugas []model.Tugas
	db.DB.Find(&tugas)
	c.JSON(http.StatusOK, gin.H{"tugas": tugas})
}

func ShowTugas(c *gin.Context) {
	var tugas []model.Tugas
	id_tugas := c.Param("id_tugas")

	if err := db.DB.First(&tugas, id_tugas).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "data tidak ditemukan"})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"tugas": tugas})
}

func CreateTugas(c *gin.Context) {
	var tugas model.Tugas

	if err := c.ShouldBindJSON(&tugas); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := dbConn.Create(&tugas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"tugas": tugas})
}

func UpdateTugas(c *gin.Context) {
	var tugas model.Tugas
	id_tugas := c.Param("id_tugas")

	if err := c.ShouldBindJSON(&tugas); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	if db.DB.Model(&tugas).Where("id_tugas=?", id_tugas).Updates(&tugas).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Tidak dapat mengedit"})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di edit"})
}

func DeleteTugas(c *gin.Context) {
	var tugas model.Tugas

	// Get tugas ID from URL parameter
	id_tugas := c.Param("id_tugas")

	// Find the tugas by ID
	if err := db.DB.First(&tugas, id_tugas).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Tugas tidak ditemukan"})
		return
	}

	// Begin transaction
	tx := db.DB.Begin()

	// Delete tugas from database
	if err := tx.Delete(&tugas).Error; err != nil {
		tx.Rollback() // Rollback transaction if there's an error
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus tugas dari database"})
		return
	}

	// Commit transaction if deletion is successful
	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
