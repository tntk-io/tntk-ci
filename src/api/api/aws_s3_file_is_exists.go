package api

import (
  "context"
  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/service/s3"
  "log"
)

func IsObjectExistsInS3Bucket(bucketName string, objectKey string) bool{
  _, err := executeHeadObjectMethodOnS3Api(bucketName, objectKey)
  return ! handleErrorOnS3HeadObjectAction(bucketName, objectKey, err)
}

func executeHeadObjectMethodOnS3Api(bucketName string, objectKey string) (output *s3.HeadObjectOutput, err error){
  s3Client := getS3Client()
  return s3Client.HeadObject(context.TODO(), &s3.HeadObjectInput{
    Bucket:               aws.String(bucketName),
    Key:                  aws.String(objectKey),
  })
}

func handleErrorOnS3HeadObjectAction(bucketName, objectKey string, err error) (hit bool){
  logUnableToHeadObjectFromS3Bucket(bucketName, objectKey, err)
  return err != nil
}

func logUnableToHeadObjectFromS3Bucket(bucketName string, objectKey string, err error) {
  if err != nil {
    log.Printf("Unable to head object {%s} from S3 bucket {%s}: { %v }", objectKey, bucketName, err)
  }
}