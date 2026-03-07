// Package blogrenderer provides functionality to render blog posts and index pages using HTML templates.
package blogrenderer

import (
	"embed"
	"html/template"
	"io"

	"github.com/gomarkdown/markdown/parser"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

// A PostRenderer is responsible for rendering blog posts and index pages using HTML templates.
type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

// NewPostRenderer creates a new PostRenderer by parsing the embedded HTML templates.
func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, mdParser: parser}, nil
}

// Render renders a single blog post to the provided io.Writer using the "blog.gohtml" template.
func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", newPostViewModel(p, r))
}

// RenderIndex renders an index page of blog posts to the provided io.Writer using the "index.gohtml" template.
func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}
