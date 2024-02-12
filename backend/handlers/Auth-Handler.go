package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"server/db"
	"server/model"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var dbConn *gorm.DB

func init() {
	// Initialize the database connection
	var err error
	dbConn, err = db.InitDB()
	if err != nil {
		// Print an error message and exit
		fmt.Println("Failed to initialize database connection:", err)
		return
	}
}

func CreateUser(c *gin.Context) {
    if dbConn == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
        return
    }

    var newUser model.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        fmt.Println(err)
        return
    }

    // Hash the password using bcrypt
    hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Passphrase), 10)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error hashing password"})
        fmt.Println(err)
        return
    }

    // Store the hashed password in the newUser object
    newUser.Passphrase = string(hash)

    // Create the user in the database
    if err := dbConn.Create(&newUser).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        fmt.Println(err)
        return
    }

    // Return success response
    c.JSON(http.StatusCreated, newUser)
}
func UpdateUser(c *gin.Context) {
	if dbConn == nil {
		// Handle the case where dbConn is nil
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection not initialized"})
		return
	}

	// Parse user ID from URL parameter
	userID := c.Param("NIS")
	var Updateuser model.User

	// Retrieve user from the database
	if err := dbConn.Where("NIS = ?", userID).First(&Updateuser).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		fmt.Println(err)
		return
	}

	// Bind JSON data to user struct
	if err := c.ShouldBindJSON(&Updateuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	  // Hash the password using bcrypt
    hash, err := bcrypt.GenerateFromPassword([]byte(Updateuser.Passphrase), 10)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error hashing password"})
        fmt.Println(err)
        return
    }
  Updateuser.Passphrase = string(hash)

	// Update user in the database
	if err := dbConn.Save(&Updateuser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, Updateuser)
}

func GetUser(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": "database connection error, or from the server error"})
		return
	}

	nis := c.Param("NIS")
	var Getuser model.User

	if err := dbConn.Where("NIS = ?", nis).First(&Getuser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user no found"})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusCreated, Getuser)

}
func DelUser(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database connection error, or from the server error"})
		return
	}

	nis := c.Param("NIS")
	var DelUser model.User
	if err := dbConn.Where("NIS = ?", nis).Delete(&DelUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user no found"})
		fmt.Println()
		return
	}
	c.JSON(http.StatusOK, gin.H{"succes": "user sucessfuly deleted"})
}

func GetALLuser(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error connection or from the server error"})
		return
	}
	var AllUser []model.User
	if err := dbConn.Find(&AllUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot find user"})
		return
	}
	c.JSON(http.StatusOK, AllUser)
}
func LoginUser(c *gin.Context) {
	var body struct {
		NIS        int    `json:"NIS"`
		Passphrase string `json:"Passphrase"`
	}
	var login model.User

	// Bind the request body to the body struct
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Retrieve the user from the database
	if err := dbConn.First(&login, "NIS = ?", body.NIS).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot find login info"})
		return
	}

	// Compare the provided password with the hashed password from the database
	if err := bcrypt.CompareHashAndPassword([]byte(login.Passphrase), []byte(body.Passphrase)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
    fmt.Println(err)
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": login.NIS,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Return the token to the client
	c.JSON(http.StatusOK, gin.H{"token": tokenString, "status": login.Status, "nama": login.Nama})
}

func GetAllUserByKelas(c *gin.Context) {
	if dbConn == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"ERROR": "database connection error, or from the server error"})
		return
	}

	kelas := c.Param("Kelas")
	var GetAlluser []model.User

	if err := dbConn.Where("Kelas = ?", kelas).Find(&GetAlluser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user no found"})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusCreated, GetAlluser)

}

func UploadProfilePicture(c *gin.Context) {
	nisParam := c.Param("NIS")
	nis, err := strconv.Atoi(nisParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid NIS"})
		return
	}

	// Retrieve the file from the form data
	file, err := c.FormFile("profilePicture")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve the file"})
		return
	}

	// Create a unique filename for the uploaded file
	fileName := filepath.Join("profile_pictures", strconv.Itoa(nis)+"_"+file.Filename)

	// Save the file to the server
	if err := c.SaveUploadedFile(file, fileName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the file"})
		return
	}

	// Update the user's Profilepicture field with the filename
	if err := dbConn.Model(&model.User{}).Where("NIS = ?", nis).Update("Profilepicture", fileName).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile picture"})
    return
}


	// Return success response
	c.JSON(http.StatusOK, gin.H{"message": "Profile picture uploaded successfully"})
}






