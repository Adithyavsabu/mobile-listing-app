package controllers

import (
	"encoding/json"
	"fmt"
	"mobile_listing_app/data"
	"net/http"
)

// Register New user
// Register New user
func UserRegistration(w http.ResponseWriter, r *http.Request) {
	var newUser data.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	newUser.Usertype = "customer"
	_, err = db.Exec("INSERT INTO appuser (username, password,usertype) VALUES ($1, $2, $3)", newUser.Username, newUser.Password, newUser.Usertype)
	if err != nil {
		fmt.Println("failed to insert into db", err)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User inserted successfully"})
}

// Validate Existing USer
func UserValidation(w http.ResponseWriter, r *http.Request) {

	// Decode the JSON request body to get the username and password

	var requestData data.LoginData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// validate the username and password

	var storedPassword string
	row := db.QueryRow("SELECT password FROM appuser WHERE username = $1", requestData.Username)
	err = row.Scan(&storedPassword)
	if err != nil {
		fmt.Println("No data Found with that username")
		return
	}

	if storedPassword != requestData.Password { // Compare the stored password with the provided password
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	} else {
		var usertype string
		query := db.QueryRow("SELECT usertype FROM appuser WHERE username = $1", requestData.Username)
		err = query.Scan(&usertype)
		if err != nil {
			fmt.Println("No usertype Found with that username")
			return
		}

		// If validation succeeds, send a success response
		jsonResponse := usertype
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jsonResponse)
	}
}
