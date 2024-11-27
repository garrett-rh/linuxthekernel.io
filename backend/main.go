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

	// kind of a hack to ensure that we serve the images properly. Will probably go away if I find it in me to add in a database or some sort of external image storage
	http.HandleFunc("GET /blog/imgs/{id}", handlers.ImageHandler)

	files := http.FileServer(http.Dir("./static"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			http.ServeFile(w, r, "./static/index.html")
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

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
