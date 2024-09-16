package main

import "net/http"

func createRoutes(router *http.ServeMux) {
	
  router.HandleFunc("/", homeHandler)
  router.HandleFunc("/features", featuresHandler)
  router.HandleFunc("/solutions", solutionsHandler)
  router.HandleFunc("/enterprise", enterpriseHandler)
  router.HandleFunc("/pricing", pricingHandler)
}

