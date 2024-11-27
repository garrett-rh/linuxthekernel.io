package helpers

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

// Post represents the metadata and content of a blog post
type PostMetadata struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Date    string   `json:"date"`
	Summary string   `json:"summary"`
	Tags    []string `json:"tags"`
}

type Post struct {
	Content string `json:"content"`
}

func GetAllPosts() ([]PostMetadata, error) {
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

func GetPostContent(id string) (Post, error) {
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
