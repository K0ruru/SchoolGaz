package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"server/auth"
	"server/db"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
		NIS        int    `json:"nis"`
		Nama       string `json:"nama"`
		Passphrase string `json:"passphrase"`
		Email      string `json:"email"`
		NoTelp     string `json:"no_telp"`
		Jenkel     string `json:"jenkel"`
		Agama      string `json:"agama"`
		Pass       string `json:"-"` 
		Token      string `json:"token"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userNis := vars["nis"]

		db, err := db.ConnectDB()
		if err != nil {
				http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
				return
		}

		defer db.Close()

		query := "SELECT nis, nama, passphrase, email, no_telp, jenkel, agama FROM users WHERE nis = $1"
		row := db.QueryRow(query, userNis)


		var user User
		err = row.Scan(&user.NIS, &user.Nama, &user.Passphrase, &user.Email, &user.NoTelp, &user.Jenkel, &user.Agama)
		if err == sql.ErrNoRows {
				http.Error(w, "User not found", http.StatusNotFound)
				return
		} else if err != nil {
				http.Error(w, "Error retrieving user", http.StatusInternalServerError)
				return
		}

		json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
				http.Error(w, "Invalid request payload", http.StatusBadRequest)
				return
		}

		// Print received user data for debugging
		fmt.Printf("Received user data: %+v\n", user)

		// Hash the user's password before storing it in the database
		hashedPassphrase, err := bcrypt.GenerateFromPassword([]byte(user.Passphrase), bcrypt.DefaultCost)
		if err != nil {
				http.Error(w, "Error hashing password", http.StatusInternalServerError)
				return
		}
		user.Passphrase = string(hashedPassphrase)

		db, err := db.ConnectDB()
		if err != nil {
				http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
				return
		}
		defer db.Close()

		query := `
    INSERT INTO users (nis, nama, passphrase, email, no_telp, jenkel, agama)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING nis;
`
err = db.QueryRow(query, user.NIS, user.Nama, user.Passphrase, user.Email, user.NoTelp, user.Jenkel, user.Agama).Scan(&user.NIS)

		if err != nil {
				http.Error(w, "Error creating user: "+err.Error(), http.StatusInternalServerError)
				return
		}

		token, err := auth.GenerateToken(user.NIS)
		if err != nil {
				http.Error(w, "Error generating token", http.StatusInternalServerError)
				return
		}
		user.Token = token

		response := struct {
				User  User   `json:"user"`
				Token string `json:"token"`
		}{
				User:  user,
				Token: token,
		}

		json.NewEncoder(w).Encode(response)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	db, err := db.ConnectDB()
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	query := "SELECT nis, passphrase FROM users WHERE nis = $1"
	row := db.QueryRow(query, user.NIS)

	var storedUser User
	err = row.Scan(&storedUser.NIS, &storedUser.Passphrase)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	// Compare the provided passphrase with the stored hashed passphrase
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Passphrase), []byte(user.Passphrase))
	if err != nil {
		http.Error(w, "Invalid passphrase", http.StatusUnauthorized)
		return
	}

	// If login is successful, generate a token
	token, err := auth.GenerateToken(user.NIS)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Respond with the token
	response := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	json.NewEncoder(w).Encode(response)
}

