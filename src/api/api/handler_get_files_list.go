package api

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

// @BasePath /api/v1

// GetFilesList godoc
// @Summary List all files available for download
// @Schemes
// @Tags example
// @Produce json
// @Success 200 {string} {}
// @Router /files [get]
func GetFilesList(ctx *gin.Context){
  handleGetFilesList(ctx)
}

func handleGetFilesList(ctx *gin.Context){
  authData := ValidateUserAuthenticationByToken(ctx)
  handleAuthorizedFilesListFromDynamoDbEntry(ctx, authData)
}

func handleAuthorizedFilesListFromDynamoDbEntry(ctx *gin.Context, authenticationData *userAuthenticationData) {
  filenamesInDynamoDbJsonString := GetFilesListAttributeFromDynamoDbEntryByUsername(authenticationData.Username)
  handleFilesListNotFoundInDynamoDbEntry(ctx, filenamesInDynamoDbJsonString)
  handleFilesListFoundInDynamoDbEntry(ctx, filenamesInDynamoDbJsonString)
}

func handleFilesListNotFoundInDynamoDbEntry(ctx *gin.Context, filenamesInDynamoDbJsonString string) {
  if isFilesListNotFoundInDynamoDbEntry(filenamesInDynamoDbJsonString) {
    ctx.String(http.StatusNotFound, "[]")
  }
}

func handleFilesListFoundInDynamoDbEntry(ctx *gin.Context, filenamesInDynamoDbJsonString string) {
  if isFilesListFoundInDynamoDbEntry(filenamesInDynamoDbJsonString) {
    ctx.String(http.StatusOK, filenamesInDynamoDbJsonString)
  }
}

func isFilesListNotFoundInDynamoDbEntry(filenamesInDynamoDbJsonString string) bool {
  return ! isFilesListFoundInDynamoDbEntry(filenamesInDynamoDbJsonString)
}

func isFilesListFoundInDynamoDbEntry(filenamesInDynamoDbJsonString string) bool{
  return filenamesInDynamoDbJsonString != ""
}