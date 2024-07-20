package asfuncss

import (
	"net/http"
	"os"
	"strings"
)

func Login(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("mydatabasefile.txt")
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)

	}
	defer file.Close()
	r.ParseForm()
	content, err := os.ReadFile("mydatabasefile.txt")
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
	}
	if len(string(content)) == 0 || string(content) == "" {
		w.Write([]byte("WRONG PASSWORD OR USERNAME"))
	} else {
		myslice := strings.Split(string(content), "\n")
		for i := 0; i < len(myslice); i++ {
			myusernameslice := strings.Split(myslice[i], " ")
			if myusernameslice[2] == "username:"+r.Form.Get("username")+"," {
				if CheckPassword(r, myusernameslice[3]) {
					w.Write([]byte("SUCCESSFULLY LOG IN"))
				} else {
					w.Write([]byte("WRONG PASSWORD OR USERNAME"))
				}
			}

		}
	}
}

func CheckPassword(r *http.Request, password string) bool {
	return "password:"+r.Form.Get("password") == password
}
