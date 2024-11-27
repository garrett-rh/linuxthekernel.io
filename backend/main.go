package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

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

		// helps with the markdown -> html conversion
		// this will likely go away if/when i move away from serving from the filesystem
		if strings.Contains(r.URL.Path, "imgs") {
			// grab the image name from the file path
			urlParts := strings.Split(r.URL.Path, "/")
			imgPath := fmt.Sprintf("./content/imgs/%s", urlParts[len(urlParts)-1])

			if _, err := os.Stat(imgPath); err == nil {
				http.ServeFile(w, r, imgPath)
			} else if os.IsNotExist(err) {
				http.NotFound(w, r)
			} else {
				http.Error(w, "Forbidden", http.StatusForbidden)
			}
			return
		}

		if _, err := os.Stat("./static" + r.URL.Path); err == nil {
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
