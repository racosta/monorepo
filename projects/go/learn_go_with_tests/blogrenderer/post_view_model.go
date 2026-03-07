package blogrenderer

import (
	"html/template"

	"github.com/gomarkdown/markdown"
)

// A postViewModel is a view model that combines a Post with its HTML-rendered body for use in templates.
type postViewModel struct {
	Post
	HTMLBody template.HTML
}

// newPostViewModel creates a new postViewModel by converting the Markdown body of the Post into HTML using the provided PostRenderer's Markdown parser.
func newPostViewModel(p Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}
