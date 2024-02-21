package api

import (
  "context"
  "fmt"
  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
  "github.com/aws/aws-sdk-go-v2/service/s3"
  "github.com/gin-gonic/gin"
)

func DownloadFileFromS3AndStoreInTemporaryFileOnFilesystem(ctx *gin.Context, username string) (pathToDownloadedFileOnFilesystem string, err error){
  temporaryFileOnFilesystem := createFile(generatePathOnFilesystemWithoutWorkspaceToDownloadedFile(getFilenameFromUri(ctx)))
  defer temporaryFileOnFilesystem.Close()

  _, err = getS3Downloader().Download(context.TODO(), temporaryFileOnFilesystem, &s3.GetObjectInput{
    Bucket: aws.String(GetS3BucketName()),
    Key:    aws.String(getS3ObjectForDownloadLocationPath(username, getFilenameFromUri(ctx))),
  })

  if err != nil {
    return "", err
  }

  return temporaryFileOnFilesystem.Name(), nil
}

func generatePathOnFilesystemWithWorkspaceToDownloadedFile(username, s3Filename string) string{
  return fmt.Sprintf("%s/%s", username, s3Filename)
}

func generatePathOnFilesystemWithoutWorkspaceToDownloadedFile(s3Filename string) string{
  return fmt.Sprintf("%s", s3Filename)
}

func getS3Downloader() *manager.Downloader{
  return manager.NewDownloader(getS3Client())
}