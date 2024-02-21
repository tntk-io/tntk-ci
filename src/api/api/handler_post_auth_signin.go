package api

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "net/http"
)

// PostAuthSignIn godoc
// @Summary Authenticate and receive token
// @Accept application/json
// @Produce json
// @Success 201 {string} {object} map[string]string
// @Router /auth/sign-in [post]
func PostAuthSignIn(ctx *gin.Context){
  handleUserAuthorization(ctx, ValidateUserAuthenticationByPassword(ctx))
}

func handleUserAuthorization(ctx *gin.Context, authenticationData *userAuthenticationData){
    handleUserAuthorized(ctx, authenticationData)
    handleUserNotAuthorized(ctx, authenticationData)
}

func handleUserAuthorized(ctx *gin.Context, authenticationData *userAuthenticationData){
  if authenticationData.IsAuthorized {
    ctx.String(http.StatusAccepted, fmt.Sprintf(`{"username": "%s", "state": "authorized", "token": "%s"}`, authenticationData.Username, authenticationData.Token))
  }
}