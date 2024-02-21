package api

import (
  "os"
  "strings"
)

func getSsmPathWithPrefix(path string) string{
  return removeDuplicatedForwardSlashes(strings.Replace(path, "ssm:", getSsmNamePrefix(), 1))
}

func getSsmNamePrefix() string{
  return os.Getenv("SSM_PREFIX")
}

func isSsmPath(rawValue string) bool{
  return strings.HasPrefix(rawValue, "ssm:/")
}