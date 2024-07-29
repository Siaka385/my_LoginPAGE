package asfuncss

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

type User struct {
	Role         string `json:"role"`
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Confpassword string `json:"confpassword"`
	Email        string `json:"email"`
}

func Reg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user := User{
		Role:         r.Form.Get("role"),
		Username:     r.Form.Get("username"),
		Password:     r.Form.Get("password"),
		Confpassword: r.Form.Get("confirmpassword"),
		Email:        r.Form.Get("email"),
	}

	// Convert ID to integer, if provided
	idStr := r.Form.Get("id")
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.Write([]byte("Invalid ID"))
			return
		}
		user.ID = id
	}

	if CheckUsernameExist(user.Username, w) {
		w.Write([]byte("Username already exists"))
	} else {
		if user.Password != user.Confpassword {
			w.Write([]byte("Password should be the same as the confirm password"))
		} else {
			user.Password = Hashpassword(user.Password)
			user.Confpassword = user.Password
			SaveDetails(user, w)
			w.Write([]byte("YOU HAVE SUCCESSFULLY SIGNED UP"))
		}
	}
}

func SaveDetails(user User, w http.ResponseWriter) {
	databaseFile := "users.json"

	var users []User
	if _, err := os.Stat(databaseFile); err == nil {
		// If the file exists, read the existing users
		fileContent, err := os.ReadFile(databaseFile)
		if err != nil {
			http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
			return
		}
		json.Unmarshal(fileContent, &users)
	}

	// Append the new user to the users slice
	users = append(users, user)

	// Marshal the updated users slice to JSON
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}

	// Write the JSON data to the file
	err = os.WriteFile(databaseFile, data, 0o644)
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
}
