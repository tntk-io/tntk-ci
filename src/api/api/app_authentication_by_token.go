package api

import "github.com/gin-gonic/gin"

func ValidateUserAuthenticationByToken(ctx *gin.Context) (authenticationData *userAuthenticationData){
  return handleUserAuthenticationByToken(ctx)
}

func handleUserAuthenticationByToken(ctx *gin.Context) (authenticationData *userAuthenticationData) {
  token := getTokenFromUri(ctx)
  user := GetUserByToken(token)

  if user != nil {
    return handleSuccessfulTokenUserAuthentication(user)
  } else {
    return handleNotSuccessfulTokenUserAuthentication(user, token)
  }

}

func handleSuccessfulTokenUserAuthentication(user *User) (authenticationData *userAuthenticationData) {
  if user != nil {
    return &userAuthenticationData{
      Username:     user.Name,
      Token:        user.Token,
      IsAuthorized: true,
    }
  }

  return nil
}

func handleNotSuccessfulTokenUserAuthentication(user *User, token string) (authenticationData *userAuthenticationData) {
  if user == nil {
    return &userAuthenticationData{
      Username:     "",
      Token:        token,
      IsAuthorized: false,
    }
  }

  return nil
}

func getTokenFromUri(ctx *gin.Context) (token string) {
  if len(ctx.Request.Header["Token"]) > 0 {
    token = ctx.Request.Header["Token"][0]
  }

  return
}