package handlers

import (
	"net/http"
	"server/database"
	"server/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IndexTugas(c *gin.Context) {
	var tugas []models.Tugas
	database.DB.Find(&tugas)
	c.JSON(http.StatusOK, gin.H{"tugas": tugas})
}

func ShowTugas(c *gin.Context) {
	var tugas []models.Tugas
	id := c.Param("id")

	if err := database.DB.First(&tugas, id).Error; err != nil {
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
	var tugas models.Tugas

	if err := c.ShouldBindJSON(&tugas); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tugas.Pembuatan = time.Now()

	database.DB.Create(&tugas)
	c.JSON(http.StatusOK, gin.H{"tugas": tugas})
}

func UpdateTugas(c *gin.Context) {
	var tugas models.Tugas
	id := c.Param("id")

	if err := c.ShouldBindJSON(&tugas); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&tugas).Where("id=?", id).Updates(&tugas).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengedit"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di edit"})
}

func DeleteTugas(c *gin.Context) {
	var tugas models.Tugas

	// Get tugas ID from URL parameter
	id := c.Param("id")

	// Find the tugas by ID
	if err := database.DB.First(&tugas, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Tugas tidak ditemukan"})
		return
	}

	// Delete tugas from database
	if err := database.DB.Delete(&tugas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus tugas dari database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
