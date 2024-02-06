package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
			fmt.Println("Error loading .env file")
	}
}

func ConnectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
	os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("pgx", connStr)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    return db, nil
}

	func InitializeDB() {
		db, err := ConnectDB()
		if err != nil {
				fmt.Println("Error connecting to the database:", err)
				return
		}
		defer db.Close()

		_, err = db.Exec(`
				CREATE TABLE IF NOT EXISTS users (
						nis SERIAL PRIMARY KEY,
						nama VARCHAR(100),
						passphrase VARCHAR(100),
						email VARCHAR(100),
						no_telp VARCHAR(20),
						jenkel VARCHAR(10),
						agama VARCHAR(20)
				);
		`)

		if err != nil {
				fmt.Println("Error creating table:", err)
		}
	}