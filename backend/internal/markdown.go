// Package markdown contains all of my markdown/post specific things
// includes things like parsing out the metadata of the files
package internal

import (
	"errors"
	"fmt"
	"github.com/yuin/goldmark/renderer/html"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

// PostMetadata represents the metadata and content of a blog post
type PostMetadata struct {
	ID      string   `json:"id"`
	Title   string   `json:"title"`
	Date    string   `json:"date"`
	Summary string   `json:"summary"`
	Tags    []string `json:"tags"`
}

// Post is just the content string after the metadata
type Post struct {
	Content string `json:"content"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}

// GetAllPosts returns a listing of all markdown files found in the content folder.
// Will then parse our the metadata and returns a list of all the posts
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

// GetPostContent returns the post content and ignores the header metadata
func GetPostContent(id string) (Post, error) {
	data, err := os.ReadFile(fmt.Sprintf("content/%s.md", id))
	if err != nil {
		return Post{}, err
	}

	var buf strings.Builder
	md := goldmark.New(
		goldmark.WithExtensions(&frontmatter.Extender{}),
		goldmark.WithRendererOptions(html.WithUnsafe()),
	)
	ctx := parser.NewContext()

	err = md.Convert(data, &buf, parser.WithContext(ctx))
	if err != nil {
		return Post{}, err
	}
	return Post{Content: buf.String()}, nil
}

// parseMarkdownMetadata extracts metadata from a markdown file
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

// extractFrontMatter parses the title from the front matter
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
		return PostMetadata{}, errors.New("failed to read metadata")
	}

	return metadata, nil
}

// SortPostsByDate takes a slice of PostMetadata and returns them sorted by date
// Under the hood it relies on slices.SortFunc() to sort
func SortPostsByDate(posts []PostMetadata) {
	slices.SortFunc(posts, func(i, j PostMetadata) int {
		iDate, err := time.Parse(time.DateOnly, i.Date)
		if err != nil {
			log.Printf("Failed to parse %s. Recieved: %s", i.Date, err.Error())
			return 0
		}
		jDate, err := time.Parse(time.DateOnly, j.Date)
		if err != nil {
			log.Printf("Failed to parse %s. Recieved: %s", j.Date, err.Error())
			return 0
		}

		return jDate.Compare(iDate)
	})
}
