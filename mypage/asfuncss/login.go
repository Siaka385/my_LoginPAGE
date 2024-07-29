package asfuncss

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// Open the JSON file containing user data
	file, err := os.Open("users.json")
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Read and unmarshal the JSON content
	var users []User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	username := r.Form.Get("loginusername")
	// Check if the username exists and the password matches
	for _, user := range users {
		if user.Username == username {
			if CheckPassword(r, user.Password) {
				w.Write([]byte("SUCCESSFULLY LOGGED IN"))
				fmt.Println("you are in")
			} else {
				w.Write([]byte("WRONG PASSWORD OR USERNAME"))
				fmt.Println(user.Password)
				fmt.Println(r.Form.Get("loginpassword"))
			}
			return
		}
	}

	// If the username was not found
	w.Write([]byte("WRONG PASSWORD OR USERNAME"))
}

func CheckPassword(r *http.Request, storedPassword string) bool {
	//compare the hashed password
	providedPassword := r.Form.Get("loginpassword")
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	return err == nil
}
