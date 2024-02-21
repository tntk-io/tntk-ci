package api

import (
  "app/api/docs"
  "github.com/gin-gonic/gin"
  "github.com/penglongli/gin-metrics/ginmetrics"
  swaggerFiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"

)

// @BasePath /api/v1

// @title DemoApp API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes https
func NewRouter() *gin.Engine{
  router := gin.New()
  router.Use(CORSMiddleware())

  router.MaxMultipartMemory = 8 << 20  // 8 MiB

  docs.SwaggerInfo.BasePath = "/api/v1"

  router.
    GET("/api/v1/file/:name", GetFileDownload).
    DELETE("/api/v1/file/:name", DeleteFile).
    POST("/api/v1/request", PostRequestForPageProcessing).
    GET("/api/v1/files", GetFilesList).
    POST("/api/v1/auth/sign-in", PostAuthSignIn).
    POST("/api/v1/auth/sign-up", PostAuthSignUp).
    GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

  m := ginmetrics.GetMonitor()
  m.SetMetricPath("/api/metrics")
  m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
  m.Use(router)

  return router
}


