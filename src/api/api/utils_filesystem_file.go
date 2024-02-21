package api

import (
  "log"
  "os"
)

func createFile(path string) *os.File{
  file, err := os.Create( path )
  if err != nil {
    log.Println(err)
  }

  return file
}