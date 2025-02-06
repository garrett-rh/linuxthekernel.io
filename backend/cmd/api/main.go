package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	c := New()
	c.PopulateConfig()
	http.HandleFunc("GET /api/posts", PostsHandler)
	http.HandleFunc("GET /api/posts/{id}", PostHandler)
	http.HandleFunc("GET /api/car_tax/localities", LocalitiesHandler)
	http.HandleFunc("POST /api/car_tax/calculate", CarTaxHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		files := http.FileServer(http.Dir("./static"))
		files.ServeHTTP(w, r)
	})

	if c.Tls {
		log.Printf("Listening on port %d", c.Port)
		log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%d", c.Port), c.Cert, c.Key, nil))
	} else {
		log.Printf("Listening on port %d", c.Port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Port), nil))
	}
}
