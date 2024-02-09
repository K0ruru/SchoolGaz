package tugasHandlers

import (
	"encoding/json"
	"net/http"
	"server/database"
	"server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var tugas []models.Tugas
	database.DB.Find(&tugas)
	c.JSON(http.StatusOK, gin.H{"tugas": tugas})
}

func Show(c *gin.Context) {
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

func Create(c *gin.Context) {
	var tugas []models.Tugas
	if err := c.ShouldBindJSON(&tugas); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	database.DB.Create(&tugas)
	c.JSON(http.StatusOK, gin.H{"tugas": tugas})

}

func Update(c *gin.Context) {
	var tugas models.Tugas
	id := c.Param("id")
	if err := c.ShouldBindJSON(&tugas); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if database.DB.Model(&tugas).Where("id=?", id).Updates(&tugas).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data tidak bisa di update"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di update"})

}

func Delete(c *gin.Context) {
	var tugas models.Tugas

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if database.DB.Delete(&tugas, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Gagal mengahpus tugas"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil di hapus"})

}
