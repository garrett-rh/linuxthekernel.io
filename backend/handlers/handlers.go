package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"linuxthekernel.io/helpers"
)

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
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the post ID from the URL path
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 4 || pathParts[1] != "api" || pathParts[2] != "posts" {
		http.NotFound(w, r)
		return
	}
	id := pathParts[3]

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
