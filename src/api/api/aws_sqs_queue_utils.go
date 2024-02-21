package api

import (
  "fmt"
  "os"
)


func getSqsQueueUrl() string {
  return getSqsQueueUrlByQueueName(getSqsQueueName())
}

func getSqsQueueUrlByQueueName(queueName string) string{
  return fmt.Sprintf("https://sqs.%s.amazonaws.com/%s/%s", getAwsRegion(), getAwsAccountId(), queueName)
}

func getSqsQueueName()string{
  return readValueFromSsmIfSsmPathProvided(os.Getenv("SQS_QUEUE_NAME"))
}