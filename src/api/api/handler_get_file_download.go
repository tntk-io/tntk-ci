package api

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "net/http"
)


// @BasePath /api/v1

// GetFileDownload godoc
// @Summary Download PDF file from server
// @Schemes
// @Description Downloads file from s3 bucket throught server
// @Tags example
// @Produce application/pdf
// @Success 200 {string} file
// @Router /file/:name [get]
func GetFileDownload(ctx *gin.Context){
  authData := ValidateUserAuthenticationByToken(ctx)
  handleAuthorizedDownloadFileFromS3AndServe(ctx, authData)
  handleUserNotAuthorized(ctx, authData)
}

func handleAuthorizedDownloadFileFromS3AndServe(ctx *gin.Context, authenticationData *userAuthenticationData){
  if authenticationData.IsAuthorized {
 //   if handleErrorOnS3BucketOrObjectMissing(ctx) { return }

    pathToTemporaryFileOnFilesystem, err := DownloadFileFromS3AndStoreInTemporaryFileOnFilesystem(ctx, authenticationData.Username)
    if err != nil {
      ctx.String(http.StatusInternalServerError, err.Error())
      return
    }

    ctx.Header("Content-Description", "File Transfer")
    ctx.Header("Content-Transfer-Encoding", "binary")
    ctx.Header("Content-Disposition", "attachment; filename="+getFilenameFromUri(ctx))
    ctx.Header("Content-Type", "application/octet-stream")
    ctx.File(pathToTemporaryFileOnFilesystem)
  }
}

func handleErrorOnS3BucketOrObjectMissing(ctx *gin.Context) (hit bool) {
  return handleErrorResponseOnS3BucketMissing(ctx) || handleErrorResponseOnObjectMissingInS3Bucket(ctx)
}

func handleErrorResponseOnS3BucketMissing(ctx *gin.Context) (hit bool){
  if hit = ! isS3BucketExists(GetS3BucketName()); hit{
    ctx.String(http.StatusNotFound, fmt.Sprintf("S3 bucket { %s } does not exist", GetS3BucketName()))
  }

  return
}

func handleErrorResponseOnObjectMissingInS3Bucket(ctx *gin.Context) (hit bool){
  if hit = ! IsObjectExistsInS3Bucket(GetS3BucketName(), getFilenameFromUri(ctx)); hit{
    ctx.String(http.StatusNotFound, fmt.Sprintf("Object { %s } is missing in S3 bucket { %s } does not exist", getFilenameFromUri(ctx), GetS3BucketName()))
  }

  return
}

func getFilenameFromUri(ctx *gin.Context) string{
  return ctx.Param("name")
}
