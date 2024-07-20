package asfuncss

import (
	"net/http"
	"os"
	"strings"
)

func Reg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("fullname")
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	confirm := r.Form.Get("confirmpassword")

	if CheckUsernameExist(username, w) {
		w.Write([]byte("Username already exist"))
	} else {
		if password != confirm {
			w.Write([]byte("password should be same with confirm password"))
		} else if strings.Contains(username, " ") {
			w.Write([]byte("Your username should not have space it"))
		} else {
			name := strings.ReplaceAll(name, " ", "+")
			SaveDetails([]string{name, username, password}, w)
			w.Write([]byte("YOU HAVE SUCCESSFUL SIGN UP"))
		}
	}
}

func SaveDetails(mydata []string, w http.ResponseWriter) {
	databasename := "mydatabasefile.txt"
	file, err := os.OpenFile(databasename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		http.Error(w, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	name := "Full name:" + mydata[0] + ", "
	username := "username:" + mydata[1] + ", "
	password := "password:" + mydata[2]
	content, _ := os.ReadFile(databasename)
	if len(string(content)) == 0 || string(content) == "" {
		file.WriteString(name + username + password)
	} else {
		file.WriteString("\n" + name + username + password)
	}
}
