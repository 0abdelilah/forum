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
