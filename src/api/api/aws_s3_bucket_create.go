package api

import (
  "context"
  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/service/s3"
  "log"
)

func CreateS3BucketIfMissing(bucketName string){
  if isS3BucketNotExists(bucketName){
    createS3Bucket(bucketName)
  }

}

func createS3Bucket(bucketName string){
  _, err := getS3Client().CreateBucket(context.TODO(), &s3.CreateBucketInput{
    Bucket: aws.String(getS3BucketNamePrefix() + bucketName),
  })

  logUnableToCreateS3BucketError(err, bucketName)
}


func getS3BucketNamePrefix() (prefix string){
  if isLocalStackModeEnabled(){
    prefix = "/"
  }
  return
}


func logUnableToCreateS3BucketError(err error, bucketName string){
  if err != nil {
    log.Printf("Unable to create S3 bucket { %s }: { %v }", bucketName, err)
  }
}