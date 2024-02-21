package api

import (
  "context"
  "github.com/aws/aws-sdk-go-v2/service/s3"
  "github.com/aws/aws-sdk-go-v2/service/s3/types"
  "github.com/thoas/go-funk"
  "log"
)

func isS3BucketNotExists(bucketName string) bool{
  return ! isS3BucketExists(bucketName)
}

func isS3BucketExists(bucketName string) bool{
  return isListOfS3BucketsOutputContainsBucket(
    getListOfS3BucketsOutput(bucketName),
    bucketName)
}

func getListOfS3BucketsOutput(bucketName string) *s3.ListBucketsOutput{
  listOfBucketsOutput, err := getS3Client().ListBuckets(context.TODO(), &s3.ListBucketsInput{})
  logUnableToListBucketsError(err)

  return listOfBucketsOutput
}

func isListOfS3BucketsOutputContainsBucket(listOfBucketsOutput *s3.ListBucketsOutput, bucketName string) bool{
  return funk.Contains(listOfBucketsOutput.Buckets, func (bucket types.Bucket) bool {
    return *bucket.Name == bucketName
  })
}


func logUnableToListBucketsError(err error) {
  if err != nil {
    log.Printf("Unable to list buckets: { %v }", err)
  }
}