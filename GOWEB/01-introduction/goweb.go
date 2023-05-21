package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Router
	// http.HandleFunc("/", Hello)
	// http.HandleFunc("/params", Params)
	// http.HandleFunc("/notfound", PageNotFound)

	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Hello)
	mux.HandleFunc("/params", Params)
	mux.HandleFunc("/notfound", PageNotFound)

	// create server
	server := &http.Server{
		Addr:    "localhost:3001",
		Handler: mux,
	}
	fmt.Println("server on: http://localhost:3001/")
	// log.Fatal(http.ListenAndServe("localhost:3001", nil))
	// log.Fatal(http.ListenAndServe("localhost:3001", mux))
	log.Fatal(server.ListenAndServe())

}

func Hello(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Mothod: " + r.Method)
	fmt.Fprintln(rw, "Hello world!")
}

func PageNotFound(rw http.ResponseWriter, r *http.Request) {
	// http.NotFound(rw, r)
	http.Error(rw, "Page not found", http.StatusNotFound)
}

func Params(rw http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.Query())
	query := r.URL.Query()
	name := query.Get("name")
	age := query.Get("age")
	fmt.Fprintf(rw, "Name: %s, age: %s\n", name, age)
}
