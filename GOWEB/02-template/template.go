package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type User struct {
	Username string
	Age      int
	Active   bool
	Admin    bool
	Courses  []Course
}

type Course struct {
	Name string
}

var templates = template.Must(template.New("T").ParseGlob("templates/*.html"))
var errors = template.Must(template.ParseFiles("templates/errors/error.html"))

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	// create server
	server := &http.Server{
		Addr:    "localhost:3001",
		Handler: mux,
	}
	fmt.Println("server on: http://localhost:3001/")
	log.Fatal(server.ListenAndServe())
}

func Index(rw http.ResponseWriter, r *http.Request) {
	// functions := template.FuncMap{
	// 	"hello": SayHello,
	// }
	// fmt.Fprintln(rw, "Hello")
	// template, err := template.New("index.html").Funcs(functions).
	// 	ParseFiles("index.html")

	courses := []Course{
		{"C1"},
		{"C2"},
		{"C3"},
	}
	user := User{"username1", 33, true, true, courses}

	//template := template.Must(template.New("template.html").ParseFiles("template.html", "base.html"))
	//template.Execute(rw, user)

	// err := templates.ExecuteTemplate(rw, "template.html", user)
	// if err != nil {
	// 	panic(err)
	// }
	renderTemplate(rw, "template.html", user)

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	template.Execute(rw, nil)
	// }
}

func SayHello(name string) string {
	return "Hello " + name + " from SayHello()"
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
