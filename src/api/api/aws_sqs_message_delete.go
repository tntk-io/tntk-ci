package api

import (
  "context"
  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/service/sqs"
  "log"
)

func DeleteSqsMessage(queueName string, receiptHandle string) (ok bool, err error){
  sqsClient := getSqsClient()
  _, err = sqsClient.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
    QueueUrl:      aws.String(getSqsQueueUrlByQueueName(queueName)),
    ReceiptHandle: aws.String(receiptHandle),
  })
  return ! logErrorOnDeleteSqsMessage(err, queueName, receiptHandle), err
}

func logErrorOnDeleteSqsMessage(err error, queueName, receiptHandle string) (isHitError bool) {
  if err != nil {
    log.Printf("Unable to delete send SQS message { %s } from queue { %s }: %v", receiptHandle, getSqsQueueUrlByQueueName(queueName), err)
  }
  return err != nil
}