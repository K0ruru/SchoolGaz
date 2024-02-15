package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"server/model"
	"strconv"

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

func UploadTugas(c *gin.Context) {
	idParam := c.Param("id_jawaban")
	id_jawaban, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid NIS"})
		return
	}

	// Retrieve the file from the form data
	file, err := c.FormFile("File")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve the file"})
		return
	}

	// Create a unique filename for the uploaded file
	fileName := filepath.Join("jawaban", strconv.Itoa(id_jawaban)+"_"+file.Filename)

	// Save the file to the server
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the file"})
		return
	}

	// Update the user's Profilepicture field with the filename
	if err := dbConn.Model(&model.User{}).Where("id_jawaban = ?", id_jawaban).Update("file", fileName).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile picture"})
		return
	}

	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "tugas  uploaded successfully"})


}



