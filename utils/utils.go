package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, text string, code int) {
	t, err := template.ParseFiles("./templates/error.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(code)
	err = t.Execute(w, struct{ Error string }{Error: text})
	if err != nil {
		fmt.Println(err)
	}
}

func ToastError(w http.ResponseWriter, text string, hash string) {
	tmpt, err := template.ParseFiles("./templates/auth.html")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}

	err = tmpt.Execute(w, struct {
		Error string
		Hash  string
	}{
		Error: text,
		Hash:  hash,
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", 500)
		return
	}
}
