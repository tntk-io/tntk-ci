package api

import (
  "encoding/json"
  "fmt"
  "github.com/gin-gonic/gin"
  "net/http"
  "time"
)


// PostRequestForPageProcessing godoc
// @Summary Post new request for processing
// @Accept application/json
// @Produce json
// @Success 201 {string} nil
// @Router /request [post]
func PostRequestForPageProcessing(ctx *gin.Context){
  handlePostRequestForPageProcessing(ctx)
}

func handlePostRequestForPageProcessing(ctx *gin.Context){
  authenticationData := ValidateUserAuthenticationByToken(ctx)
  handleAuthorizedPostRequestForPageProcessing(ctx, authenticationData)
  handleUserNotAuthorized(ctx, authenticationData)
}

func handleAuthorizedPostRequestForPageProcessing(ctx *gin.Context, authenticationData *userAuthenticationData){
  if authenticationData.IsAuthorized {
    ctx.Header("Content-Type", "application/json")
    if ok, body := handleReadBody(ctx); ok {
      if ok, request := handleJsonUnmarshalRequestForPageProcessing(ctx, body); ok {
        enrichSqsRequestWithUsernameAttribute(request, authenticationData)
        enrichSqsRequestWithTimestampAttribute(request, authenticationData)
        handleSendSqsMessageForPageProcessingRequest(ctx, request)
      }
    }
  }
}

func handleJsonUnmarshalRequestForPageProcessing(ctx *gin.Context, body []byte) (ok bool, request *requestForPageProcessing){
  err := json.Unmarshal(body, &request)
  return ! handleFailedPostRequestForPageProcessingOnJsonUnmarshalError(ctx, err), request
}

func enrichSqsRequestWithUsernameAttribute(request *requestForPageProcessing, authenticationData *userAuthenticationData){
  request.Username = authenticationData.Username
}
func enrichSqsRequestWithTimestampAttribute(request *requestForPageProcessing, authenticationData *userAuthenticationData){
  request.Username = authenticationData.Username
  request.Timestamp = time.Now().Unix()
}

func handleSendSqsMessageForPageProcessingRequest(ctx *gin.Context, request *requestForPageProcessing){
  _, err := sendSqsMessage(getSqsQueueName(), request.toJsonString())
  handleFailedSendSqsMessageForPageProcessingRequest(ctx, err)
  handleSuccessfulSendSqsMessageForPageProcessingRequest(ctx, err)
}
func handleFailedSendSqsMessageForPageProcessingRequest(ctx *gin.Context, err error){
  if err != nil {
    ctx.String(http.StatusInternalServerError, fmt.Sprintf("Error sending SQS message for page processing request: %v", err))
  }
}

func handleSuccessfulSendSqsMessageForPageProcessingRequest(ctx *gin.Context, err error){
  if err == nil {
    ctx.String(http.StatusCreated, "")
  }
}

func handleFailedPostRequestForPageProcessingOnJsonUnmarshalError(ctx *gin.Context, err error) (isErrorHit bool){
  if err != nil {
    ctx.String(http.StatusMethodNotAllowed, fmt.Sprintf("Error handling request body content: %v", err))
    return true
  }
  return false
}