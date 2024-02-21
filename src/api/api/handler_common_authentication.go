package api

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "net/http"
)

func handleUserNotAuthorized(ctx *gin.Context, authenticationData *userAuthenticationData) (){
  if ! authenticationData.IsAuthorized {
    ctx.Header("Content-Type", "application/json")
    ctx.String(http.StatusUnauthorized, fmt.Sprintf(`{"username": "%s", "state": "unauthorized", "token": ""}`, authenticationData.Username))
  }
}