package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"linuxthekernel.io/helpers"
)

// takes the requested /blog/imgs/{id} path and transforms that to ./content/imgs/{id}
// the reason for this is because the path is being modified by the react router.
// this will probably die once this thing develops a little more
func ImageHandler(w http.ResponseWriter, r *http.Request) {

	imgPath := fmt.Sprintf("./content/imgs/%s", r.PathValue("id"))

	if _, err := os.Stat(imgPath); err == nil {
		http.ServeFile(w, r, imgPath)
	} else if os.IsNotExist(err) {
		http.NotFound(w, r)
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

// returns the listing of all posts
func PostsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests allowed", http.StatusMethodNotAllowed)
		return
	}

	posts, err := helpers.GetAllPosts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// handler for fetching a single post by ID
func PostHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	content, err := helpers.GetPostContent(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
