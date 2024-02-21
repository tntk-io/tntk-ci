package api

import (
  "context"
  "github.com/aws/aws-sdk-go-v2/service/sqs"
  "github.com/thoas/go-funk"
  "log"
)

func isSqsQueueNotExists(queueName string) bool{
  return ! isSqsQueueExists(queueName)
}

func isSqsQueueExists(queueName string) bool{
  return isListOfSqsQueuesOutputContainsQueueByQueueName(
    getListOfSqsQueuesOutput(),
    queueName)
}

func getListOfSqsQueuesOutput() *sqs.ListQueuesOutput{
  listOfQueuesOutput, err := getSqsClient().ListQueues(context.TODO(), &sqs.ListQueuesInput{})
  logUnableToListSqsQueuesError(err)

  return listOfQueuesOutput
}

func isListOfSqsQueuesOutputContainsQueueByQueueName(listOfQueuesOutput *sqs.ListQueuesOutput, queueName string) bool{
  return funk.Contains(listOfQueuesOutput.QueueUrls, func (queueUrl string) bool {
    return queueUrl == getSqsQueueUrlByQueueName(queueName)
  })
}


func logUnableToListSqsQueuesError(err error) {
  if err != nil {
    log.Printf("Unable to list buckets: { %v }", err)
  }
}