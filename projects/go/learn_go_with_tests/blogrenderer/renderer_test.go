package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/racosta/monorepo/projects/go/learn_go_with_tests/blogrenderer"
	"github.com/sebdah/goldie/v2"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		g := goldie.New(t)
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		g.Assert(t, "rendered-post", buf.Bytes())
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		g := goldie.New(t)
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		g.Assert(t, "rendered-index", buf.Bytes())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogrenderer.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		if err := postRenderer.Render(io.Discard, aPost); err != nil {
			b.Fatal(err)
		}
	}
}
