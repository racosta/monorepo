// Package blogposts provides a way to read blog posts from a file system and return them as a slice of Post structs.
package blogposts

import (
	"errors"
	"io"
	"io/fs"
)

// NewPostsFromFS reads all the files in the given file system and returns a slice of Post structs representing the
// blog posts contained in those files. Each file is expected to have a specific format, with the title, description,
// tags, and body of the post separated by specific lines. If there is an error reading any of the files, it returns
// an error.
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (p Post, err error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}

	defer closeWithErr(postFile, &err)

	return newPost(postFile)
}

func closeWithErr(c io.Closer, err *error) {
	closeErr := c.Close()
	if closeErr != nil {
		if *err == nil {
			*err = closeErr
		} else {
			*err = errors.Join(*err, closeErr)
		}
	}
}
