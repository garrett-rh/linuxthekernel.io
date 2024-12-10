// Package handlers contains API http handlers
package handlers

import (
	"encoding/json"
	"net/http"

	"linuxthekernel.io/markdown"
)

// PostsHandler returns the listing of all posts
func PostsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests allowed", http.StatusMethodNotAllowed)
		return
	}

	posts, err := markdown.GetAllPosts()
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

// PostHandler handler for fetching a single post by ID
func PostHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	content, err := markdown.GetPostContent(id)
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
