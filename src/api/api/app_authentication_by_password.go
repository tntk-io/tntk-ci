package api

import "github.com/gin-gonic/gin"

func ValidateUserAuthenticationByPassword(ctx *gin.Context) (authorizationData *userAuthenticationData){
  username, password, ok := ctx.Request.BasicAuth()
  if ok {
      return handleUserBasicAuthenticationOnSuccessfulPasswordMatch(username, password)
  }

  return handleUserBasicAuthenticationOnFailure(username)
}

func handleUserBasicAuthenticationOnSuccessfulPasswordMatch(username, password string) (authorizationData *userAuthenticationData){
  user := GetUserByName(username)
  if user != nil {
    if user.IsPasswordMatch(password) {
      token := getUuid()
      err := UpdateUserToken(user.Name, token)
      if err == nil {
        return newUserAuthorizationByPasswordDataOnUserAuthorized(user.Name, token)
      }
    }
  }

  return nil
}

func handleUserBasicAuthenticationOnFailure(username string) (authorizationData *userAuthenticationData){
  return newUserAuthorizationByPasswordDataOnUserNotAuthorized(username)
}


func newUserAuthorizationByPasswordDataOnUserAuthorized(username, token string) *userAuthenticationData {
  return &userAuthenticationData{
    Username:     username,
    Token:        token,
    IsAuthorized: true,
  }
}

func newUserAuthorizationByPasswordDataOnUserNotAuthorized(username string) *userAuthenticationData {
  return &userAuthenticationData{
    Username:     username,
    Token:        getUuid(),
    IsAuthorized: false,
  }
}