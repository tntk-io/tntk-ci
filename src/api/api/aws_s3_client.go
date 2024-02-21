package api

import "github.com/aws/aws-sdk-go-v2/service/s3"

var s3Client *s3.Client

func getS3Client() *s3.Client{
  if s3Client == nil {
    s3Client = s3.NewFromConfig(NewAwsConfig())
  }

  return s3Client
}
