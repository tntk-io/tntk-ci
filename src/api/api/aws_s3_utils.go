package api

import (
  "fmt"
  "os"
  "strings"
)

func getS3ObjectForDownloadLocationPath(workspace, filename string) (objectPath string){
  //objectPath = fmt.Sprintf("%s%s", getS3BucketDirectoryPathForProcessedFiles(), filename)
  objectPath = fmt.Sprintf("%s/%s", workspace, filename)
  return getS3PathCleanedOfMultipleSlashes(objectPath)
}

func getS3BucketDirectoryPathForProcessedFiles() (path string){
  path = getS3PathForProcessingResults()
  if isLocalStackModeEnabled(){
    path = fmt.Sprintf("%s/%s", GetS3BucketName(), getS3PathForProcessingResults())
  }
  return getS3PathCleanedOfMultipleSlashes(path)
}

func getS3BucketUploadPath() (path string) {
  path = getS3PathForFilesToProcess()
  if isLocalStackModeEnabled(){
    path = fmt.Sprintf("%s/%s", GetS3BucketName(), getS3PathForFilesToProcess())
  }
  return getS3PathCleanedOfMultipleSlashes(path)
}

func GetS3BucketName() string {
  return readValueFromSsmIfSsmPathProvided(os.Getenv("S3_BUCKET_NAME"))
}

func getS3PathCleanedOfMultipleSlashes(path string) string{
  buffer := strings.ReplaceAll(path, "///", "/")
  return strings.ReplaceAll(buffer, "//", "/")
}

func getS3PathForFilesToProcess() string{
  return os.Getenv("S3_PATH_FOR_FILES_TO_PROCESS")
}

func getS3PathForProcessingResults() string{
  return os.Getenv("S3_PATH_FOR_FILES_PROCESSING_RESULTS")
}