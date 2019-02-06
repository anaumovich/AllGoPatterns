package main

import (
	"AllGoPatterns/ProjectOfRegistrationForm/controller"
	"fmt"
	"net/http"
	"strings"
)

func handler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := strings.Join([]string{r.Method, r.URL.Path}, "")
		x := controller(key)
		x(w, r)
	}
}

func controller(key string) func(w http.ResponseWriter, r *http.Request) {
	switch key {
	case "GET/":
		return Controller.StartController()
	case "GET/login":
		return Controller.LoginController()
	case "GET/registration":
		return Controller.RegistrationController()
	default:
		return Controller.ErrorController()
	}
}

func main() {
	http.HandleFunc("/", handler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
