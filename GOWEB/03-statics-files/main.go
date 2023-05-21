package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type User struct {
	Username string
	Age      int
	Active   bool
	Admin    bool
}

var templates = template.Must(template.New("T").ParseGlob("templates/*.html"))
var errors = template.Must(template.ParseFiles("templates/errors/error.html"))

func main() {
	staticFile := http.FileServer(http.Dir("static"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	// Mux Static File
	mux.Handle("/static/", http.StripPrefix("/static/", staticFile))

	// create server
	server := &http.Server{
		Addr:    "localhost:3001",
		Handler: mux,
	}
	fmt.Println("server on: http://localhost:3001/")
	log.Fatal(server.ListenAndServe())
}

func Index(rw http.ResponseWriter, r *http.Request) {
	user := User{"username1", 33, true, true}
	renderTemplate(rw, "index.html", user)
}

func register() {

}

func renderTemplate(rw http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(rw, name, data)
	if err != nil {
		//panic(err)
		//http.Error(rw, "Internal Server Error !!", http.StatusInternalServerError)
		handleError(rw, http.StatusInternalServerError)
	}
}

func handleError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	errors.Execute(rw, nil)
}
