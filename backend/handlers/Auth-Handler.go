package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	auth "server/Auth"
	"server/db"
	"strconv"
	"strings"

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
    Profilepicture sql.NullString `json:"profilepicture"`
    Status     string `json:"status"`
    Role       string `json:"role"`
    Kelas      int    `json:"kelas"`
    Mapel      int    `json:"mapel"`
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
    _, err = dbConn.Exec("INSERT INTO users (nis, nama, passphrase, email, gender, agama) VALUES ($1, $2, $3, $4, $5, $6)",
        newUser.NIS, newUser.Name, hashedPassword, newUser.Email, newUser.Gender, newUser.Religion)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user",})
        fmt.Println(err)
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

    // Mendapatkan data entitas yang ingin diperbarui dari body permintaan
    var updateData map[string]interface{}
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        fmt.Println(err)
        return
    }

    // Memeriksa apakah entitas yang ingin diperbarui ada dalam daftar entitas yang diperbolehkan
    allowedEntities := map[string]bool{"name": true, "passphrase": true, "email": true, "gender": true, "religion": true, "profilepicture": true, "status": true, "role": true, "kelas": true, "mapel": true}
    for key := range updateData {
        if _, ok := allowedEntities[key]; !ok {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid entity to update"})
            return
        }
    }

    // Jika passphrase diperbarui, hash passphrase baru
    if newPassword, ok := updateData["passphrase"]; ok {
        newPass, ok := newPassword.(string)
        if !ok {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid passphrase format"})
            return
        }
        if newPass != "" {
            hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPass), bcrypt.DefaultCost)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
                return
            }
            updateData["passphrase"] = string(hashedPassword)
        }
    }

    // Menyiapkan query pembaruan
    query := "UPDATE users SET"
    var values []interface{}
    i := 1
    for key, value := range updateData {
        query += " " + key + "=$" + strconv.Itoa(i) + ","
        values = append(values, value)
        i++
    }
    // Menghapus koma ekstra di akhir query
    query = strings.TrimSuffix(query, ",")

    // Menambahkan kondisi WHERE
    query += " WHERE nis=$" + strconv.Itoa(i)
    values = append(values, nis)

    // Melakukan operasi pembaruan pengguna
    _, err = dbConn.Exec(query, values...)
    if err != nil {
        fmt.Println("Error updating user:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
        fmt.Println(err)
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
// LoginHandler handles user login
func LoginHandler(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Retrieve user from the database by NIS
    query := "SELECT nis, nama, passphrase, email, gender, agama, status, role FROM users WHERE nis = $1"
    row := dbConn.QueryRow(query, req.NIS)

    // Initialize a User struct to store the result
    var user User
    err := row.Scan(&user.NIS, &user.Name, &user.Passphrase, &user.Email, &user.Gender, &user.Religion, &user.Status, &user.Role)
    if err != nil {
        // Check if the user is not found
        if err == sql.ErrNoRows {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid NIS or password"})
            return
        }
        // Handle other errors
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
        return
    }

    // Check if user exists and verify password
    if !checkPassword(req.Passphrase, user.Passphrase) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid NIS or password"})
        return
    }

    // Generate JWT token
    token, err := auth.CreateToken(req.NIS)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    // Return token and user details to the client
    c.JSON(http.StatusOK, gin.H{"token": token, "status": user.Status, "name": user.Name, "role": user.Role})
}


// checkPassword compares the provided password with the hashed password from the database
func checkPassword(password string, hashedPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}


func GetSiswaFromKelas(c *gin.Context) {
    kelasStr := c.Param("kelas")

    // cek id kelasnya
    if kelasStr == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Kelas ID is required"})
        return
    }

    // ubah id nya tadi jadi int
    kelas, err := strconv.Atoi(kelasStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Kelas ID format"})
        fmt.Println(err)
        return
    }

    // querry untuk ngambil semua siswa yang kelasnya sama
    rows, err := dbConn.Query("SELECT nis, nama, email, profilepicture, gender FROM users WHERE kelas=$1", kelas)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
        fmt.Println(err)
        return
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.NIS, &user.Name, &user.Email, &user.Profilepicture, &user.Gender); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan users"})
            fmt.Println(err)
            return
        }
        users = append(users, user)
    }

    if err := rows.Err(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
        fmt.Println(err)
        return
    }

    c.JSON(http.StatusOK, users)
}