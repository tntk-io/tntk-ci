package api

import (
  "crypto/md5"
  "fmt"
)

//func GetMd5FromString(string string) string{
//  data := []byte(string)
//  return fmt.Sprintf("%x", md5.Sum(data))
//}

func GetMd5FromString(string string) string{
  return fmt.Sprintf("%x", md5.Sum([]byte(string)))
}