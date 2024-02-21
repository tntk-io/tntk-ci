package api

import (
  "github.com/aws/aws-sdk-go-v2/service/sqs"
)

var sqsClient *sqs.Client

func getSqsClient() *sqs.Client{
  if sqsClient == nil {
    sqsClient = sqs.NewFromConfig(NewAwsConfig())
  }

  return sqsClient
}

