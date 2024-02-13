package handlers

import (
	"fmt"
	"net/http"
	"os"
	"server/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func CreateGuru(c *gin.Context) {
	if dbConn == nil {
		// Handle the case where dbConn is nil
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
		return
	}

	var newGuru model.Guru
	if err := c.ShouldBindJSON(&newGuru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	// Hash the password using bcrypt
    hash, err := bcrypt.GenerateFromPassword([]byte(newGuru.Passphrase), 10)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error hashing password"})
        fmt.Println(err)
        return
    }

    // Store the hashed password in the newUser object
    newGuru.Passphrase = string(hash)
	if err := dbConn.Create(&newGuru).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusCreated, newGuru)
}
func Loginguru(c *gin.Context) {
    var body struct {
        NIS        int    `json:"NIS"`
        Passphrase string `json:"Passphrase"`
    }
    var login model.Guru

    // Bind the request body to the body struct
    if err := c.Bind(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    // Retrieve the user from the database
    if err := dbConn.First(&login, "NIS = ?", body.NIS).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot find login info"})
        fmt.Println(err)
        return
    }

    // Compare the provided password with the hashed password from the database
    if err := bcrypt.CompareHashAndPassword([]byte(login.Passphrase), []byte(body.Passphrase)); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
        fmt.Println(err)
        return
    }

    // Check user role
    if login.Role != "admin" && login.Role != "guru" {
        c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized role"})
        return
    }

    // Create JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub":  login.NIS,
        "role": login.Role, // Include user role in the token claims
        "exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
    })

    // Sign the token with the secret key
    tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    // Return the token and role to the client
    c.JSON(http.StatusOK, gin.H{"token": tokenString, "role": login.Role})
}


func Updateguru(c *gin.Context) {
	if dbConn == nil {
		// Handle the case where dbConn is nil
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
		return
	}

	// Parse user ID from URL parameter
	guruID := c.Param("NIS")
	var Updateguru model.Guru

	// Retrieve user from the database
	if err := dbConn.Where("NIS = ?", guruID).First(&Updateguru).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		fmt.Println(err)
		return
	}

	// Bind JSON data to user struct
	if err := c.ShouldBindJSON(&Updateguru); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	  // Hash the password using bcrypt
    hash, err := bcrypt.GenerateFromPassword([]byte(Updateguru.Passphrase), 10)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error hashing password"})
        fmt.Println(err)
        return
    }
  Updateguru.Passphrase = string(hash)

	// Update user in the database
	if err := dbConn.Save(&Updateguru).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, Updateguru)
}

func GetGuru(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": "database connection error, or from the server error"})
		return
	}

	nis := c.Param("NIS")
	var GetGuru model.Guru

	if err := dbConn.Preload("MapelData").Where("NIS = ?", nis).First(&GetGuru).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "guru no found"})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusCreated, GetGuru)

}
func DelGuru(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection error, or from the server error"})
		return
	}

	nis := c.Param("NIS")
	var DelGuru model.Guru
	if err := dbConn.Where("NIS = ?", nis).Delete(&DelGuru).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user no found"})
		fmt.Println()
		return
	}
	c.JSON(http.StatusOK, gin.H{"succes": "guru sucessfuly deleted"})
}

func GetALLguru(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connection or from the server error"})
		return
	}
  var AllGuru []model.Guru
	if err := dbConn.Preload("MapelData").Find(&AllGuru).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find guru"})
		return
	}
	c.JSON(http.StatusOK, AllGuru)
}

