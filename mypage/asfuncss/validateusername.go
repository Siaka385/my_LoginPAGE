package asfuncss

import (
	"net/http"
	"os"
	"strings"
)

func CheckUsernameExist(m string, w http.ResponseWriter) bool {
	file, err := os.Open("mydatabasefile.txt")
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return false
	}
	defer file.Close()

	content, err := os.ReadFile("mydatabasefile.txt")
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return false
	}
	if len(string(content)) == 0 || string(content) == "" {
		return false
	} else {
		myslice := strings.Split(string(content), "\n")
		for i := 0; i < len(myslice); i++ {
			myusernameslice := strings.Split(myslice[i], " ")
			if myusernameslice[2] == "username:"+m+"," {
				return true
			}
		}

	}
	return false
}
