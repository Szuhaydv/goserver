package main

import (
  "log"
  "net/http"
  "fmt"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a request at my domain")
	w.Write([]byte("Hello, welcome to my project!"))
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
