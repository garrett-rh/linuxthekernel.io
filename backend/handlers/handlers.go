package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

// Post represents the metadata and content of a blog post
type PostMetadata struct {
	ID    string   `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Tags  []string `json:"tags"`
}

type Post struct {
	Content string `json:"content"`
}

func getAllPosts() ([]PostMetadata, error) {
	var posts []PostMetadata
	files, err := filepath.Glob("content/*.md")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		post, err := parseMarkdownMetadata(file)
		if err == nil {
			posts = append(posts, post)
		}
	}
	return posts, nil
}

func getPostContent(id string) (Post, error) {
	data, err := os.ReadFile(fmt.Sprintf("content/%s.md", id))
	if err != nil {
		return Post{}, err
	}

	var buf strings.Builder
	md := goldmark.New(
		goldmark.WithExtensions(&frontmatter.Extender{}),
	)
	ctx := parser.NewContext()

	err = md.Convert(data, &buf, parser.WithContext(ctx))
	if err != nil {
		return Post{}, err
	}
	return Post{Content: buf.String()}, nil
}

// extracts metadata and content from a markdown file
func parseMarkdownMetadata(filename string) (PostMetadata, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return PostMetadata{}, err
	}

	metadata, err := extractFrontMatter(data)
	if err != nil {
		return PostMetadata{}, err
	}

	return metadata, nil
}

// parses the title from the front matter (or default if not found)
func extractFrontMatter(data []byte) (PostMetadata, error) {
	var buf strings.Builder
	var metadata PostMetadata
	md := goldmark.New(
		goldmark.WithExtensions(&frontmatter.Extender{}),
	)
	ctx := parser.NewContext()

	err := md.Convert(data, &buf, parser.WithContext(ctx))
	if err != nil {
		return PostMetadata{}, err
	}
	d := frontmatter.Get(ctx)
	if err := d.Decode(&metadata); err != nil {
		return PostMetadata{}, errors.New("Failed to read metadata")
	}

	return metadata, nil
}

// returns the listing of all posts
func PostsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests allowed", http.StatusMethodNotAllowed)
		return
	}

	posts, err := getAllPosts()
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

	content, err := getPostContent(id)
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
