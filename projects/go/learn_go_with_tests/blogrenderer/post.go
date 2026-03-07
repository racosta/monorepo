package blogrenderer

import "strings"

// A Post represents a blog post with a title, description, tags, and body.
type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

// SanitizedTitle returns a URL-friendly version of the post's title.
func (p Post) SanitizedTitle() string {
	return strings.ToLower(strings.ReplaceAll(p.Title, " ", "-"))
}
