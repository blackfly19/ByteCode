package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there I love %s!", r.URL.Path[1:])
}

func usingTemplates(w http.ResponseWriter, r *http.Request) {
	home := Page{"First Page", "Hey hello"}
	temp, err := template.ParseFiles("templates/dummy.html")
	fmt.Println(err)
	temp.Execute(w, home)
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/template/", usingTemplates)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
