// Package blogrenderer provides functionality to render blog posts and index pages using HTML templates.
package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

// A PostRenderer is responsible for rendering blog posts and index pages using HTML templates.
type PostRenderer struct {
	templ *template.Template
}

// NewPostRenderer creates a new PostRenderer by parsing the embedded HTML templates.
func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

// Render renders a single blog post to the provided io.Writer using the "blog.gohtml" template.
func (r *PostRenderer) Render(w io.Writer, p Post) error {
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}

	return nil
}

// RenderIndex renders an index page of blog posts to the provided io.Writer using the "index.gohtml" template.
func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	if err := r.templ.ExecuteTemplate(w, "index.gohtml", posts); err != nil {
		return err
	}

	return nil
}
