package blogposts_test

import (
  blogposts "github.com/turnerem/blogposts"
  "testing"
  "testing/fstest"
  "io/fs"
  "errors"
  "reflect"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
  return nil, errors.New("stub fs failed")
}

func TestNewBlogPosts(t *testing.T) {
  t.Run("all posts successfully extracted from fs", func(t *testing.T) {
    fs := fstest.MapFS{
      "hey.md": {Data: []byte("Title: Post 1")},
      "ya.md": {Data: []byte("Title: Post 2")},
    }

    posts, err := blogposts.NewPostsFromFS(fs)

    if err != nil {
      t.Fatal(err)
    }

    assertPostLength(t, posts, fs)

    got := posts[0]
    want := blogposts.Post{Title: "Post 1"}

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %+v want %+v", got, want)
    }
  })

  t.Run("errors out is fs fails", func(t *testing.T) {
    _, err := blogposts.NewPostsFromFS(StubFailingFS{})

    if err == nil {
      t.Errorf("Expected an error")
    }
  })
}

func assertPostLength(t *testing.T, got []blogposts.Post, fs fstest.MapFS) {
  t.Helper()

  if len(got) != len(fs) {
    t.Errorf("got %d posts, want %d posts", len(got), len(fs))
  }
}
