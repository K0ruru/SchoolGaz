package handlers

import (
	"fmt"
	"net/http"
	"server/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllTugas(c *gin.Context) {
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

	id_tugas := c.Param("id_tugas")
	var Gettugas model.Tugas

	if err := dbConn.Where("id_tugas = ?", id_tugas).First(&Gettugas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tugas no found"})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusCreated, Gettugas)

}

func CreateTugas(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": "database connection error, or from the server error"})
		return
	}
	var tugas model.Tugas

	if err := c.ShouldBindJSON(&tugas); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tugas.CreateAt = time.Now()

	if err := dbConn.Create(&tugas).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, tugas)
}

func UpdateTugas(c *gin.Context) {

	var tugas model.Tugas
	id_tugas := c.Param("id_tugas")

	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
		return
	}

	_, err := strconv.Atoi(id_tugas)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_tugas"})
		return
	}

	if err := dbConn.Where("id_tugas = $1", id_tugas).First(&tugas).Error; err != nil {
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

	id_tugas := c.Param("id_tugas")

	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
		return
	}

	_, err := strconv.Atoi(id_tugas)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id_tugas"})
		return
	}

	if err := dbConn.Where("id_tugas = $1", id_tugas).Delete(&tugas).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"success": "tugas berhasil di delete"})
}
