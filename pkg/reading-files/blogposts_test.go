package blogposts_test

import (
	"testing"
	"testing/fstest"

	blogposts "github.com/korber710/go-with-tests/pkg/reading-files"
	"github.com/stretchr/testify/assert"
)

func TestNewBlogPosts(t *testing.T) {
	t.Parallel()

	fs := fstest.MapFS{
		"new-blog.md":     {Data: []byte(`This is my new blog.`)},
		"another-blog.md": {Data: []byte(`This is another blog.`)},
	}

	posts := blogposts.NewBlogPosts(fs)

	assert.Equal(t, len(posts), len(fs))
}
