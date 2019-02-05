package main

import (
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
	case "GET/list":
		return firstController()
	case "GET/add":
		return secondController()
	case "GET/":
		return defaultController()
	default:
		return errorController()
	}
}

func firstController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(printDefaultPage("firstController")))
		fmt.Println(err)
	}
}

func secondController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(printDefaultPage("secondController")))
		fmt.Println(err)
	}
}

func defaultController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(printDefaultPage("defaultController")))
		fmt.Println(err)
	}
}

func errorController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(printDefaultPage("fatal error")))
		fmt.Println(err)
	}
}

func printDefaultPage(child ...string) string {
	return fmt.Sprint(`<!DOCTYPE html>
		<html>
			<head>
				<title>Page Title</title>
			</head>
		<body>

		<p>`, child, `</p>

		</body>
		</html>`)
}

func main() {
	http.HandleFunc("/", handler())
	fmt.Println(" starting server on localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
