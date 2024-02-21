package api

import (
  "context"
  "fmt"
  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/config"
)

func NewAwsConfig() aws.Config{

  var cfg aws.Config

  if isLocalStackModeEnabled() {
    customResolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
      return aws.Endpoint{
        PartitionID:   "aws",
        URL:           "http://localhost:4566",
        SigningRegion: "us-east-1",
      }, nil
      return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
    })

    cfg, err = config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolver(customResolver))
  } else {
    cfg, err = config.LoadDefaultConfig(context.TODO())
  }

  if err != nil {
    panic("configuration error, " + err.Error())
  }

  return cfg
}