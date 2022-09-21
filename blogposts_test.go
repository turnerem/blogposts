package blogposts_test

import (
  blogposts "github.com/turnerem/blogposts"
  "testing"
  "testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
  fs := fstest.MapFS{
    "hey.md": {Data: []byte("hi")},
    "ya.md": {Data: []byte("ya")},
  }

  posts := blogposts.NewPostsFromFs(fs)

  if len(posts) != len(fs) {
    t.Errorf("got %d posts, want %d posts", len(posts), len(fs))
  }
}