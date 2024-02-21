package api

import (
  "github.com/gin-gonic/gin"
  "os"
)


func PostFileUpload(ctx *gin.Context) {
  if authenticationData := ValidateUserAuthenticationByToken(ctx); authenticationData.IsAuthorized {
    statusCode, responseBody := uploadFilesToS3(ctx, authenticationData)
    updateEntryInDynamoDb(authenticationData.Username, responseBody, statusCode)
    ctx.String(statusCode, responseBody)
  }
}

func isLocalStackModeEnabled() bool{
  return os.Getenv("LOCALSTACK_MODE_ENABLED") == "true"
}
