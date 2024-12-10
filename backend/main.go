package main

import (
	"errors"
	"fmt"
	"io/fs"
	"linuxthekernel.io/config"
	"log"
	"net/http"
	"os"

	"linuxthekernel.io/handlers"
)

func main() {
	c := config.New()
	c.PopulateConfig()
	http.HandleFunc("GET /api/posts", handlers.PostsHandler)
	http.HandleFunc("GET /api/posts/{id}", handlers.PostHandler)

	files := http.FileServer(http.Dir("./static"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			files.ServeHTTP(w, r)
			return
		}

		if _, err := os.Stat("./static" + r.URL.Path); err == nil {
			files.ServeHTTP(w, r)
		} else if errors.Is(err, fs.ErrNotExist) {
			http.ServeFile(w, r, "./static/index.html")
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})

	if c.Tls {
		log.Printf("Listening on port %d", c.Port)
		log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%d", c.Port), c.Cert, c.Key, nil))
	} else {
		log.Printf("Listening on port %d", c.Port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Port), nil))
	}
}
