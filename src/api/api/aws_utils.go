package api

import "os"

func getAwsRegion() string{
  return os.Getenv("AWS_REGION")
}