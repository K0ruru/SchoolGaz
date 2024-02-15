// db.go

package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)
var DB *gorm.DB
func init() {
    if err := godotenv.Load(); err != nil {
        fmt.Println("Error loading .env file")
    }
}
//koneksi database
func Connection() (*gorm.DB, error) {
    // Retrieving environment variables
    DB_USER := os.Getenv("DB_USER")
    DB_PASSWORD := os.Getenv("DB_PASSWORD")
    DB_NAME := os.Getenv("DB_NAME")
    DB_HOST := os.Getenv("DB_HOST")
    DB_PORT := os.Getenv("DB_PORT")

    // Constructing connection string
    CS := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

   db, err := gorm.Open("postgres", CS)
    if err != nil {
        return nil, err
    }
    
    // Perform a database operation to verify the connection
    if err := db.DB().Ping(); err != nil {
        db.Close() 
        return nil, err
    }
    
    return db, nil
}


func InitDB() (*gorm.DB, error) {
    db, err := Connection()
    if err != nil {
        return nil, err
    }
    return db, nil
}

