package handlers

import (
	"fmt"
	"net/http"
	"server/model"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetJawaban(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connection or from the server error"})
		return
	}
	var jawaban []model.Jawaban
	if err := dbConn.Find(&jawaban).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find jawaban"})
		return
	}
	c.JSON(http.StatusOK, jawaban)
}

func GetJwabanByid(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": "database connection error, or from the server error"})
		return
	}
	id_jawaban := c.Param("id_jawaban")
	var jawaban model.Jawaban

	if err := dbConn.Where("id_jawaban = ?", id_jawaban).First(&jawaban).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "jawaban not found"})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, &jawaban)

}

func CreateJawaban(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": "database connection error, or from the server error"})
		return
	}
	var jawaban model.Jawaban

	if err := c.ShouldBindJSON(&jawaban); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mengambil NIS dari sesi
	session := sessions.Default(c)
	nis := session.Get("NIS") // Ubah "nis" menjadi "NIS" sesuai dengan sesi yang telah Anda set sebelumnya
	if nis == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NIS not found in session"})
		return
	}

	// Konversi nis ke tipe int
	nisInt, ok := nis.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NIS in session is not an integer"})
		return
	}

	jawaban.SiswaNIS = nisInt

	jawaban.CreateAt = time.Now()

	if err := dbConn.Create(&jawaban).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, jawaban)
}

func UpdateJawaban(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
		return
	}

	id_jawaban := c.Param("id_jawaban")

	_, err := strconv.Atoi(id_jawaban)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_jawaban"})
		return
	}
	var jawaban model.Jawaban

	if err := dbConn.Where("id_jawaban = $1", id_jawaban).First(&jawaban).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	if err := c.ShouldBindJSON(&jawaban); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := dbConn.Save(&jawaban).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}

	c.JSON(http.StatusOK, jawaban)
}

func DeleteJawaban(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
		return
	}

	id_jawaban := c.Param("id_jawaban")

	_, err := strconv.Atoi(id_jawaban)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_jawaban"})
		return
	}
	var jawaban model.Jawaban

	if err := dbConn.Where("id_jawaban = $1", id_jawaban).Delete(&jawaban).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "jawaban deleted"})
}
