package api

import (
  "io"
  "io/ioutil"
  "mime/multipart"
  "strings"
)

func ProcessFile(file multipart.File) io.Reader{
  fileContent, err := ioutil.ReadAll(file)
  if err != nil {
    panic(err)
  }
  return strings.NewReader("!!!processeed!!!" + string(fileContent))
}