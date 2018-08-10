package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var files, _ = filepath.Glob("templates/*")
var templates, _ = template.New("").ParseFiles(files...)

func main() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", struct{}{})
}
