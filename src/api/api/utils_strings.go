package api

import "strings"

func removeDuplicatedForwardSlashes(raw string) (result string){
  for strings.Contains(raw, "//"){
    raw = strings.ReplaceAll(raw, "//", "/")
  }

  return raw
}