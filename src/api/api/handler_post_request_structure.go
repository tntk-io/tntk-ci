package api

import (
  "encoding/json"
  "log"
  "net/url"
)

type requestForPageProcessing struct {
  Url string
  Username string
  Timestamp int64
}

func (request requestForPageProcessing) isContainsValidUrl() bool{
  u, err := url.Parse(request.Url)
  return err == nil && u.Scheme != "" && u.Host != ""
}

func (request requestForPageProcessing) toJsonString() (jsonString string){
  jsonStringBytes, err := json.Marshal(request)
  if logErrorOnJsonMarshalRequestForPageProcessingStructureObject(err){ return }
  return string(jsonStringBytes)
}

func logErrorOnJsonMarshalRequestForPageProcessingStructureObject(err error) (isErrorHit bool){
  if err != nil {
    log.Printf("Unable to marshal RequestForPageProcessing structure object to JSON: %v", err)
  }
  return err != nil
}

func UnmarshalFromJsonRequestForPageProcessing(jsonString string) (request *requestForPageProcessing, err error){
  err = json.Unmarshal([]byte(jsonString), &request)
  return
}