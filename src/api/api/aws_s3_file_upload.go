package api

import (
  "context"
  "encoding/json"
  "fmt"
  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
  "github.com/aws/aws-sdk-go-v2/service/s3"
  "github.com/gin-gonic/gin"
  "io"
  "log"
  "mime/multipart"
  "net/http"
  "os"
  "path/filepath"
  "time"
)

func uploadFilesToS3(ctx *gin.Context, authenticationData *userAuthenticationData) (statusCode int, responseBody string){

  form, _ := ctx.MultipartForm()
  files := form.File["upload[]"]

  var filesForUpload = map[string]multipart.File{}

  for _, file := range files {
    filesForUpload[file.Filename], _ = file.Open()
  }

  uploadedFileNames := map[string]string{}

  for fileName, file := range filesForUpload {

    s3FileName := fmt.Sprintf("%d", time.Now().Unix())
    uploadedFileNames[fileName] = s3FileName

    var s3UploadPath = getS3PathCleanedOfMultipleSlashes(fmt.Sprintf("%s/%s/%s%s", getS3BucketUploadPath(), authenticationData.Username, s3FileName, filepath.Ext(fileName)))

    UploadFileToS3BucketPath(GetS3BucketName(), s3UploadPath, file)
  }

  json, _ := json.Marshal(uploadedFileNames)
  responseBody = string(json)
  statusCode = http.StatusCreated
  return
}

func UploadFileToS3BucketPath(s3BucketName string, s3UploadPath string, file io.Reader) (ok bool, err error){
  _, err = manager.NewUploader(getS3Client()).Upload(context.TODO(), &s3.PutObjectInput{
    Bucket: aws.String(GetS3BucketName()),
    Key:    aws.String(s3UploadPath),
    Body:   file,
  })

  return logFailedUploadFileToS3BucketPath(err, s3BucketName, s3UploadPath)
}

func UploadFileFromLocalStorageToS3Bucket(s3BucketName string, s3UploadPath string, localFilePath string) (ok bool, err error){

  file, err := os.Open(localFilePath)
  if isErrorHit, _ := logFailedOpenLocalFile(err, localFilePath); isErrorHit{ return false, err }

  _, err = manager.NewUploader(getS3Client()).Upload(context.TODO(), &s3.PutObjectInput{
    Bucket: aws.String(s3BucketName),
    Key:    aws.String(s3UploadPath),
    Body:   file,
  })

  return logFailedUploadFileToS3BucketPath(err, s3BucketName, s3UploadPath)
}

func logFailedOpenLocalFile(err error, filepath string)(isErrorHit bool, e error){
  if err != nil {
    log.Printf("Unable to open file { %s } for read: %v", filepath, err)
  }
  return err != nil, err
}

func logFailedUploadFileToS3BucketPath(err error, s3BucketName, s3UploadPath string)(isErrorNotHit bool, e error){
  if err != nil {
    log.Printf("Unable to upload file to s3 bucket {s3://%s/%s}: %v", s3BucketName, s3UploadPath, err)
  }
  return err == nil, err
}