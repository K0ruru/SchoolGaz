// db.go

package db

import (
    "database/sql"
    "fmt"
    "os"

    "github.com/joho/godotenv"
)

func init() {
    if err := godotenv.Load(); err != nil {
        fmt.Println("Error loading .env file")
    }
}

func Connection() (*sql.DB, error) {
    DB_USER := os.Getenv("DB_USER")
    DB_PASSWORD := os.Getenv("DB_PASSWORD")
    DB_NAME := os.Getenv("DB_NAME")


    CS := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)

    db, err := sql.Open("postgres", CS)
    if err != nil {
        return nil, err
    }
    err = db.Ping()
    if err != nil {
        return nil, err
    }
    return db, nil
}

func InitDB() (*sql.DB, error) {
    db, err := Connection()
    if err != nil {
        return nil, err
    }
    return db, nil
}

