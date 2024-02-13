package handlers

import (
	"fmt"
	"net/http"
	"server/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateMapel(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erorr database connection or from the server error"})
		return
	}

	var newMapel model.Mapel
	if err := c.ShouldBindJSON(&newMapel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := dbConn.Create(&newMapel).Error; err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, newMapel)

}

func UpdateMapel(c *gin.Context) {
	if dbConn == nil {
		// Handle the case where dbConn is nil
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
		return
	}

	// Parse model ID from URL parameter
	userID := c.Param("id_mapel")
	var UpdateMapel model.Mapel
  
  _, err := strconv.Atoi(userID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid kelasID"})
        return
    }


	// Retrieve model from the database
	if err := dbConn.Where("id_mapel = ?", userID).First(&UpdateMapel).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "mapel not found"})
		fmt.Println(err)
		return
	}

	// Bind JSON data to user struct
	if err := c.ShouldBindJSON(&UpdateMapel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

  	// Update model in the database
	if err := dbConn.Save(&UpdateMapel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update mapel"})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, UpdateMapel)
}



