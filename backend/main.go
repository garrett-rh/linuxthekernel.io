package main

import (
	"log"
	"net/http"
	"os"

	"linuxthekernel.io/handlers"
)

func main() {
	http.HandleFunc("/api/posts", handlers.PostsHandler)
	http.HandleFunc("/api/posts/", handlers.PostHandler)

	fs := http.FileServer(http.Dir("./static"))

	// Custom handler to support client-side routing
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			http.ServeFile(w, r, "./static/index.html")
			return
		}

		_, err := os.Stat("./static" + r.URL.Path)
		if err == nil {
			fs.ServeHTTP(w, r)
		} else if os.IsNotExist(err) {
			http.ServeFile(w, r, "./static/index.html")
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
