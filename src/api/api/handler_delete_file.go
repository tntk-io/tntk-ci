package api

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

// @BasePath /api/v1

// DeleteFile godoc
// @Summary Delete file from s3 bucket
// @Schemes
// @Produce text
// @Success 200 {string} {}
// @Router /file/:name [delete]

func DeleteFile(ctx *gin.Context){
  authData := ValidateUserAuthenticationByToken(ctx)
  handleAuthorizedDeleteFileFromS3(ctx, authData)
  handleUserNotAuthorized(ctx, authData)
}

func handleAuthorizedDeleteFileFromS3(ctx *gin.Context, authenticationData *userAuthenticationData){
  if authenticationData.IsAuthorized {
    if ! IsObjectExistsInS3Bucket(GetS3BucketName(), getS3ObjectForDownloadLocationPath(authenticationData.Username, getFilenameFromUri(ctx))) { return }
    ok, err := DeleteFileFromS3Bucket(GetS3BucketName(), getS3ObjectForDownloadLocationPath(authenticationData.Username, getFilenameFromUri(ctx)))
    handleSuccessfulFileDeletionFromS3(ctx, ok, authenticationData)
    handleUnsuccessfulFileDeletionFromS3(ctx, err)
  }
}

func handleSuccessfulFileDeletionFromS3(ctx *gin.Context, ok bool, authenticationData *userAuthenticationData){
  if ok {
    removeFilenameFromEntryInDynamoDb(authenticationData.Username, getFilenameFromUri(ctx))
    ctx.Header("Content-Description", "File removal request result")
    ctx.Header("Content-Type", "text/plain")
    ctx.String(http.StatusOK, "")
  }
}

func handleUnsuccessfulFileDeletionFromS3(ctx *gin.Context, err error){
  if err != nil {
    ctx.Header("Content-Description", "File removal request result")
    ctx.Header("Content-Type", "text/plain")
    ctx.String(http.StatusInternalServerError, err.Error())
  }
}
