package main

import (
  "app/api"
  "context"
  "encoding/json"
  "errors"
  "fmt"
  . "github.com/SebastiaanKlippert/go-wkhtmltopdf"
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
  "log"
  "strings"
  "time"
)

func HandleRequest(ctx context.Context, sqsEvent events.SQSEvent) error {
  if len(sqsEvent.Records) == 0 {
    return errors.New("No SQS message passed to function")
  }


  for _, msg := range sqsEvent.Records {
    fmt.Printf("SQS message: %s\n", msg.Body)
    request, err := api.UnmarshalFromJsonRequestForPageProcessing(msg.Body)
    queueArnItems := strings.Split(msg.EventSourceARN, ":")
    queueName := queueArnItems[len(queueArnItems) - 1]
    api.DeleteSqsMessage(queueName, msg.ReceiptHandle)
    if err == nil {
      wkhtml(request.Url, request.Username)
    } else {
      fmt.Printf("Error unmarshalling SQS body: %v\n", err)
    }
  }

  return nil
}

//func main(){
//  wkhtml("https://smile.io", "parsec")
//}

func main() {
 lambda.Start(HandleRequest)
}


func wkhtml(url string, workspace string) {
  // Create new PDF generator
  pdfg, err := NewPDFGenerator()
  if err != nil {
    log.Fatal(err)
  }

  // Set global options
  pdfg.Dpi.Set(300)
  pdfg.Orientation.Set(OrientationLandscape)
  pdfg.Grayscale.Set(true)

  // Create a new input page from an URL
  //page := NewPage("https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf")
  page := NewPage(url)
  // Set options for this page
  page.FooterRight.Set("[page]")
  page.FooterFontSize.Set(10)
  page.Zoom.Set(0.95)

  // Add to document
  pdfg.AddPage(page)

  // Create PDF document in internal buffer
  err = pdfg.Create()
  if err != nil {
    log.Fatal(err)
  }

  // Write buffer contents to file on disk

  //temporaryFileName :=

  s3FileName := fmt.Sprintf("%s.%d.pdf", api.GetMd5FromString(url), time.Now().Unix())
  localFilePath := fmt.Sprintf("/tmp/%s", s3FileName)

  err = pdfg.WriteFile(localFilePath)
  if err != nil {
    log.Fatalf("Unable to save file { %s } on local storage: %v\n", localFilePath, err)
  }

  ok, _ := api.UploadFileFromLocalStorageToS3Bucket(api.GetS3BucketName(), fmt.Sprintf("%s/%s", workspace, s3FileName), localFilePath)
  s3ObjectKeyFullPath := fmt.Sprintf("s3://%s/%s/%s", api.GetS3BucketName(), workspace, s3FileName)
  if ok{
    fmt.Printf("Success: local file { %s } uploaded to S3 bucket { %s }\n", localFilePath, s3ObjectKeyFullPath)
    var existingListOfFilesInDynamoDb []string
    jsonString := api.GetFilesListAttributeFromDynamoDbEntryByUsername(workspace)
    if jsonString != ""{
      err := json.Unmarshal([]byte(jsonString), &existingListOfFilesInDynamoDb)
      if err != nil {
        log.Fatalf("Unable to unmarshal json string { %s } of existing list of files in dynamodb: %v\n", jsonString, err)
      } else {
        existingListOfFilesInDynamoDb = append(existingListOfFilesInDynamoDb, s3FileName)
      }
    }else {
      existingListOfFilesInDynamoDb = append(existingListOfFilesInDynamoDb, s3FileName)
    }
    api.CreateEntryInDynamoDB(workspace, existingListOfFilesInDynamoDb)
  } else {
    log.Fatalf("Failure: local file { %s } not uploaded to S3 bucket { %s }\n", localFilePath, s3ObjectKeyFullPath)
  }
}