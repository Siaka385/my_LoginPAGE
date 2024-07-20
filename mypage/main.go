package main

import (
	"fmt"
	"net/http"
	"text/template"

	"mypage/asfuncss"
)

func SignUphandler(w http.ResponseWriter) {
	tmp, _ := template.ParseFiles("signup.html")

	tmp.Execute(w, nil)
}

func Loginpageload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tmp, _ := template.ParseFiles("login.html")
	tmp.Execute(w, nil)
}

func router(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		SignUphandler(w)
	} else if r.URL.Path == "/reg" {
		asfuncss.Reg(w, r)
	} else if r.URL.Path == "/log" {
		Loginpageload(w, r)
	} else if r.URL.Path == "/login" {
		asfuncss.Login(w, r)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", router)
	fmt.Println("RUNNING SERVER")
	http.ListenAndServe("localhost:9000", mux)
}
