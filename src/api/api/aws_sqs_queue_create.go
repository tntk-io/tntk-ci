package api

import (
  "context"
  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/service/sqs"
  "log"
)

func CreateSqsQueueIfMissing(queueName string){
  if isSqsQueueNotExists(queueName){
    createSqsQueue(queueName)
  }

}

func createSqsQueue(queueName string){
  _, err := getSqsClient().CreateQueue(context.TODO(), &sqs.CreateQueueInput{
    QueueName:  aws.String(queueName),
    Attributes: map[string]string{"VisibilityTimeout":"361", "FifoQueue":"true", "ContentBasedDeduplication":"true"},
    Tags:       nil,
  })

  logUnableToCreateSqsQueueError(err, queueName)
}

func logUnableToCreateSqsQueueError(err error, queueName string){
  if err != nil {
    log.Printf("Unable to create SQS queue { %s }: { %v }", queueName, err)
  }
}