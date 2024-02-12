package handlers

import (
	"net/http"
	"server/db"
	"server/model"
	"time"

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
	id := c.Param("id")

	if err := db.DB.First(&tugas, id).Error; err != nil {
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tugas.Pembuatan = time.Now()

	db.DB.Create(&tugas)
	c.JSON(http.StatusOK, gin.H{"tugas": tugas})
}

func UpdateTugas(c *gin.Context) {
	var tugas model.Tugas
	id := c.Param("id")

	if err := c.ShouldBindJSON(&tugas); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if db.DB.Model(&tugas).Where("id=?", id).Updates(&tugas).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengedit"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di edit"})
}

func DeleteTugas(c *gin.Context) {
	var tugas model.Tugas

	// Get tugas ID from URL parameter
	id := c.Param("id")

	// Find the tugas by ID
	if err := db.DB.First(&tugas, id).Error; err != nil {
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
