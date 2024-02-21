package api

import (
  "github.com/aws/aws-sdk-go-v2/service/ssm"
)

var ssmClient *ssm.Client

func getSsmClient() *ssm.Client{
  if ssmClient == nil {
    ssmClient = ssm.NewFromConfig(NewAwsConfig())
  }

  return ssmClient
}
