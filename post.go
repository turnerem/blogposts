package blogposts

import (
  "io"
  "bufio"
  "strings"
)

// Post is...
type Post struct {
  Title		string
  Description	string
}

const (
  titleSeparator	= "Title: "
  descSeparator		= "Description: "
)

func newPost(postFile io.Reader) (Post, error) {
  scanner := bufio.NewScanner(postFile)

  readLine := func(tagName string) string {
    scanner.Scan()
    return strings.TrimPrefix(scanner.Text(), tagName)
  }

  title:= readLine(titleSeparator)
  description:= readLine(descSeparator)

  return Post{
    Title: title,
    Description: description,
  }, nil
}
