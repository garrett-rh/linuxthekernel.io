package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

func main() {
	c := New()
	c.PopulateConfig()
	files := http.FileServer(http.Dir("./static"))
	http.HandleFunc("GET /api/posts", PostsHandler)
	http.HandleFunc("GET /api/posts/{id}", PostHandler)
	http.HandleFunc("GET /api/car_tax/localities", LocalitiesHandler)
	http.HandleFunc("POST /api/car_tax/calculate", CarTaxHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			files.ServeHTTP(w, r)
			return
		}

		if _, err := os.Stat("./static" + r.URL.Path); err == nil {
			// if the path exists in our static dir, we should serve it
			files.ServeHTTP(w, r)
		} else if errors.Is(err, fs.ErrNotExist) {
			// If the file doesn't exist it is likely a react-based path.
			// Going to just serve the index and react-router should handle it
			http.ServeFile(w, r, "./static/index.html")
		} else {
			// redirecting to root if we manage to fall through to this
			http.Redirect(w, r, "/", http.StatusFound)
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
