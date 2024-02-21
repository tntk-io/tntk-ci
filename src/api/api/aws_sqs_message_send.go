package api

import (
  "context"
  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/service/sqs"
  "log"
)

func sendSqsMessage(queueName string, messageBody string) (ok bool, err error){
  sqsClient := getSqsClient()

  qUrl := getSqsQueueUrlByQueueName(queueName)

  _, err = sqsClient.SendMessage(context.TODO(), &sqs.SendMessageInput{
    MessageBody:             aws.String(messageBody),
    MessageGroupId:          aws.String("default"),
    QueueUrl:                aws.String(qUrl),
  })
  return ! logErrorOnSendSqsMessage(err, queueName), err
}

func logErrorOnSendSqsMessage(err error, queueName string) (isHitError bool) {
  if err != nil {
    log.Printf("Unable to send SQS message to queue { %s }: %v", getSqsQueueUrlByQueueName(queueName), err)
  }
  return err != nil
}