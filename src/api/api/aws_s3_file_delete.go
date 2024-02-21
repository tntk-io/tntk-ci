package api

import (
  "context"
  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/service/s3"
  "log"
)

func DeleteFileFromS3Bucket(bucketName string, objectKey string) (ok bool, err error){
  _, err = executeDeleteObjectMethodOnS3Api(bucketName, objectKey)
  return ! handleErrorOnDeleteObjectMethodOfS3Api(bucketName, objectKey, err), err
}

func executeDeleteObjectMethodOnS3Api(bucketName string, objectKey string) (output *s3.DeleteObjectOutput, err error){
  s3Client := getS3Client()
  return s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
    Bucket:               aws.String(bucketName),
    Key:                  aws.String(objectKey),
  })
}

func handleErrorOnDeleteObjectMethodOfS3Api(bucketName, objectKey string, err error) (hit bool){
  logUnableToDeleteObjectFromS3Bucket(bucketName, objectKey, err)
  return err != nil
}

func logUnableToDeleteObjectFromS3Bucket(bucketName string, objectKey string, err error) {
  if err != nil {
    log.Printf("Unable to delete object {%s} from S3 bucket {%s}: { %v }", objectKey, bucketName, err)
  }
}