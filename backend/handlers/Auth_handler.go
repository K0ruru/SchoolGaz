package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"server/db"
	"strconv"
  "server/Auth"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
    NIS        int    `json:"nis"`
    Name       string `json:"name"`
    Passphrase string `json:"passphrase"`
    Email      string `json:"email"`
    Gender     string `json:"gender"`
    Religion   string `json:"religion"`
    Status     string `json:"status"`
}

var dbConn *sql.DB


func init() {
    // Initialize the database connection
    var err error
    dbConn, err = db.InitDB()
    if err != nil {
        // Handle error
        panic(err)
    }
}

// Create user handler
func CreateUser(c *gin.Context) {
    var newUser User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Passphrase), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }

    // Gunakan password yang di-hash saat memasukkan data ke dalam database
    _, err = dbConn.Exec("INSERT INTO users (nis, nama, passphrase, email, gender, agama, status) VALUES ($1, $2, $3, $4, $5, $6, $7)",
        newUser.NIS, newUser.Name, hashedPassword, newUser.Email, newUser.Gender, newUser.Religion, newUser.Status)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user",})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
// Update user handler
func UpdateUser(c *gin.Context) {
    // Mendapatkan ID pengguna dari parameter
    nisStr := c.Param("nis")

    // Memeriksa apakah nilai "nis" ada
    if nisStr == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
        return
    }

    // Memeriksa apakah nilai "nis" adalah bilangan bulat
    nis, err := strconv.Atoi(nisStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
        return
    }

    // Mendapatkan data pengguna yang ingin diperbarui dari body permintaan
    var updateUser User
    if err := c.ShouldBindJSON(&updateUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    // Lakukan operasi pembaruan pengguna sesuai dengan kebutuhan aplikasi Anda
    _, err = dbConn.Exec("UPDATE users SET nama=$1, passphrase=$2, email=$3, gender=$4, agama=$5, status=$6 WHERE nis=$7",
        updateUser.Name, updateUser.Passphrase, updateUser.Email, updateUser.Gender, updateUser.Religion, updateUser.Status, nis)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }

    // Respon berhasil
    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}


// Get all users handler
func GetUsers(c *gin.Context) {
    rows, err := dbConn.Query("SELECT nis, nama, email, gender, agama, status FROM users")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
        return
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.NIS, &user.Name, &user.Email, &user.Gender, &user.Religion, &user.Status); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan users"})
            return
        }
        users = append(users, user)
    }

    if err := rows.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan fetch users"})
        return
    }

    c.JSON(http.StatusOK, users)
}

// GetUserByNIS retrieves a user from the database by NIS
func GetUserByNIS(nis int) (*User, error) {
    // Query the database to retrieve the user with the specified NIS
    query := "SELECT nis, nama, passphrase, email, gender, agama, status FROM users WHERE nis = $1"
    row := dbConn.QueryRow(query, nis)

    // Initialize a User struct to store the result
    var user User
    err := row.Scan(&user.NIS, &user.Name, &user.Passphrase, &user.Email, &user.Gender, &user.Religion, &user.Status)
    if err != nil {
        // Check if the user is not found
        if err == sql.ErrNoRows {
            return nil, nil
        }
        // Handle other errors
        return nil, err
    }

    // Return the user
    return &user, nil
}

// DeleteUser handles the deletion of a user
func DeleteUser(c *gin.Context) {
    nisStr := c.Param("nis")
    
    // Memeriksa apakah nilai "nis" ada
    if nisStr == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
        return
    }

    // Memeriksa apakah nilai "nis" adalah bilangan bulat
    nis, err := strconv.Atoi(nisStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
        return
    }
   _, err = dbConn.Exec("DELETE FROM users WHERE nis=$1", nis)
      if err != nil {
         fmt.Println("Failed to delete user: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}



 // LoginRequest represents the JSON structure of login request
type LoginRequest struct {
    NIS        int    `json:"nis" binding:"required"`
    Passphrase string `json:"passphrase" binding:"required"`
}

// LoginHandler handles user login
func LoginHandler(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Retrieve user from the database by NIS
    user, err := GetUserByNIS(req.NIS)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
        return
    }

    // Check if user exists and verify password
    if user == nil || !checkPassword(req.Passphrase, user.Passphrase) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid NIS or password"})
        return
    }

    // Generate JWT token
    token, err := auth.CreateToken(req.NIS)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    // Return token to the client
    c.JSON(http.StatusOK, gin.H{"token": token})
}

// checkPassword compares the provided password with the hashed password from the database
func checkPassword(password string, hashedPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}
