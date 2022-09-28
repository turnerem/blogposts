package blogposts

import (
  "io/fs"
)

// NewPostsFromFS is...
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
  dir, err := fs.ReadDir(fileSystem, ".")

  if err != nil {
    return nil, err
  }

  var posts []Post

  for _, f := range dir {
    post, err := getPost(fileSystem, f.Name())

    if err != nil {
      return nil, err //TODO: clarify err
    }

    posts = append(posts, post)
  }

  return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
  // open file
  postFile, err := fileSystem.Open(fileName)

  if err != nil {
    return Post{}, err
  }

  defer postFile.Close()

  post, err := newPost(postFile)

  // close file

  return post, err
}
