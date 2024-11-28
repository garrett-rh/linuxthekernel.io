package main

import (
	"errors"
	"io/fs"
	"log"
	"net/http"
	"os"

	"linuxthekernel.io/handlers"
)

func main() {
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

	log.Println("Starting server on :443")
	log.Fatal(http.ListenAndServeTLS(":443", "/secrets/chain.pem", "/secrets/priv.key", nil))
}
