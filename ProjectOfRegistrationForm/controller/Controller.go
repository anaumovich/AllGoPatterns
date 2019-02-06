package Controller

import (
	"AllGoPatterns/ProjectOfRegistrationForm/View"
	"fmt"
	"net/http"
)

func StartController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(View.PrintStartPage("Зарегистрируйтесь или авторизуйтесь в системе")))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ErrorController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(View.PrintDefaultPage("fatal error")))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func LoginController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(View.PrintLoginPage("Вход")))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func RegistrationController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(View.PrintDefaultPage("Регистрация в системе")))
		if err != nil {
			fmt.Println(err)
		}
	}
}
