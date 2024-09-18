package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type PageVariables struct {
  Heading string
  Content string
}

var tmpl1 = template.Must(template.ParseFiles("home.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a request at my domain")

  data := PageVariables{
    Heading: "Welcome to the Home Page",
    Content: "This is a dynamically generated HTML response",
  }

  tmpl1.ExecuteTemplate(w, "home.html", data)
}

func featuresHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to Features")
}

func solutionsHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to Solutions")
}

func enterpriseHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to Enterprise")
}

func pricingHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to Pricing")
}
