package api

import (
  "context"
  "github.com/aws/aws-sdk-go-v2/service/sts"
  "log"
)

var stsClient *sts.Client

func getStsClient() *sts.Client{
  if stsClient == nil {
    stsClient = sts.NewFromConfig(NewAwsConfig())
  }
  return stsClient
}

func getAwsAccountId() string {
  return *getStsCallerIdentity().Account
}

func getStsCallerIdentity() *sts.GetCallerIdentityOutput{
  stsCallerIdentity, err := getStsClient().GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
  if err != nil {
    log.Printf("Unable to retrieve STS Caller Identity: %v", err)
    return nil
  }
  return stsCallerIdentity
}