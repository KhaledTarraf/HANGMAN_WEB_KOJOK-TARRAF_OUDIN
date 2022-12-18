package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/KhaledTarraf/hangman-classic/src"
)

var HangmandataWeb = src.HangManData{}

type PageData struct {
	Username    string
	Difficulty  string
	HangmanData src.HangManData
}

func Homepage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := PageData{
		Username:   r.FormValue("username"),
		Difficulty: r.FormValue("difficulty"),
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Hangmanpage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("hangman.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	username := r.Form.Get("username")
	difficulty := r.Form.Get("difficulty")
	hangmanData := src.HangManData{Attempts: HangmandataWeb.Attempts, ToFind: HangmandataWeb.ToFind}
	data := PageData{
		Username:    username,
		Difficulty:  difficulty,
		HangmanData: hangmanData,
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", Homepage)
	http.HandleFunc("/hangman", Hangmanpage)

	fmt.Println("(http://localhost:8080/) - Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
