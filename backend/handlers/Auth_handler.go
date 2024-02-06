package handlers

import (
    "database/sql"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "server/db"
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

    _, err := dbConn.Exec("INSERT INTO users (nis, name, passphrase, email, gender, religion, status) VALUES ($1, $2, $3, $4, $5, $6, $7)",
        newUser.NIS, newUser.Name, newUser.Passphrase, newUser.Email, newUser.Gender, newUser.Religion, newUser.Status)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// Get all users handler
func GetUsers(c *gin.Context) {
    rows, err := dbConn.Query("SELECT nis, name, email, gender, religion, status FROM users")
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
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
        return
    }

    c.JSON(http.StatusOK, users)
}

// Update user handler
func UpdateUser(c *gin.Context) {
    nisStr := c.Param("nis")
    nis, err := strconv.Atoi(nisStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    var updatedUser User
    if err := c.ShouldBindJSON(&updatedUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err = dbConn.Exec("UPDATE users SET name=$1, passphrase=$2, email=$3, gender=$4, religion=$5, status=$6 WHERE nis=$7",
        updatedUser.Name, updatedUser.Passphrase, updatedUser.Email, updatedUser.Gender, updatedUser.Religion, updatedUser.Status, nis)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// Delete user handler
func DeleteUser(c *gin.Context) {
    nisStr := c.Param("nis")
    nis, err := strconv.Atoi(nisStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    _, err = dbConn.Exec("DELETE FROM users WHERE nis=$1", nis)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}


 
