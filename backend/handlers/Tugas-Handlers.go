package handlers

import (
	"fmt"
	"net/http"
	"server/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexTugas(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connection or from the server error"})
		return
	}
	var tugas []model.Tugas
	if err := dbConn.Find(&tugas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find tugas"})
		return
	}
	c.JSON(http.StatusOK, tugas)
}

func GetTugas(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": "database connection error, or from the server error"})
		return
	}

	id := c.Param("id")
	var Gettugas model.Tugas

	if err := dbConn.Where("id = ?", id).First(&Gettugas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tugas no found"})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusCreated, Gettugas)

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

	c.JSON(http.StatusOK, tugas)
}

func UpdateTugas(c *gin.Context) {

	var tugas model.Tugas
	id := c.Param("id")

	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
		return
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	if err := dbConn.Where("id = $1", id).First(&tugas).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	if err := c.ShouldBindJSON(&tugas); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if err := dbConn.Save(&tugas).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil di edit"})
}

func DeleteTugas(c *gin.Context) {
	var tugas model.Tugas

	id := c.Param("id")

	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
		return
	}

	_, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	if err := dbConn.Where("id = $1", id).Delete(&tugas).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"success": "tugas berhasil di delete"})
}
