package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var files, _ = filepath.Glob("templates/*")
var templates, _ = template.New("").ParseFiles(files...)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", struct{}{})
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))

	log.Println(templates)

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
