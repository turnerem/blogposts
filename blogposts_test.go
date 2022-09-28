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
  const (
    firstBody = `Title: Post 1
Description: Description 1`
    secondBody = `Title: Post 2
Description: Description 2`
  )

  fs := fstest.MapFS{
    "hey.md": {Data: []byte(firstBody)},
    "ya.md": {Data: []byte(secondBody)},
  }

  posts, err := blogposts.NewPostsFromFS(fs)

  if err != nil {
    t.Fatal(err)
  }

  assertPostLength(t, posts, fs)

  assertPost(t, posts[0], blogposts.Post{
    Title:		"Post 1",
    Description:	"Description 1",
  })
}

func assertPostLength(t *testing.T, got []blogposts.Post, fs fstest.MapFS) {
  t.Helper()

  if len(got) != len(fs) {
    t.Errorf("got %d posts, want %d posts", len(got), len(fs))
  }
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
  t.Helper()

  if !reflect.DeepEqual(got, want) {
    t.Errorf("got %+v want %+v", got, want)
  }
}
