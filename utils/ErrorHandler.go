package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, text string, code int) {
	if code != 404 {
		fmt.Println("Error:", text)
	}

	t, err := template.ParseFiles("./templates/error.html")
	if err != nil {
		http.Error(w, text, code)
		return
	}

	w.WriteHeader(code)
	t.Execute(w, struct{ ErrorText string }{ErrorText: text})
}

func ToastError(w http.ResponseWriter, text string, hash string) {
	tmpt, err := template.ParseFiles("./templates/auth.html")
	if err != nil {
		ErrorHandler(w, "Internal server error", 500)
		return
	}

	err = tmpt.Execute(w, struct {
		Text string
		Hash string
	}{
		Text: text,
		Hash: hash,
	})
	if err != nil {
		ErrorHandler(w, "Internal server error", 500)
		return
	}
}
