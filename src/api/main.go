package main

import (
  "app/api"
  "log"
  "os"
)

func main(){
  log.SetOutput(os.Stdout)
  api.Initialize()
  api.NewRouter().Run(os.Getenv("API_SOCKET"))
}


