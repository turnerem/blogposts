package blogpost

// Post is...
type Post struct {
  Title string
}

func newPost(postFile io.Reader) (Post, error) {
  // read all of file
  postData, err := io.ReadAll(postFile)
  if err != nil {
    return Post{}, err
  }

  // pop it in the Post struct
  post := Post{Title: string(postData)[7:]}

  return post, nil
}
