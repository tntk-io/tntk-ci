package api

import (
  "context"
  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/service/ssm"
  "log"
)

func readValueFromSsmIfSsmPathProvided(value string) (finalValue string){
  if isSsmPath(value){
    value = readFromSsm(value)
  }
  return value
}

func readFromSsm(path string) (value string){
  parameter, err := getSsmClient().GetParameter(context.TODO(), &ssm.GetParameterInput{
    Name:           aws.String(getSsmPathWithPrefix(path)),
    WithDecryption: true,
  })

  logUnableToReadFromSsmError(err, getSsmPathWithPrefix(path))

  return *parameter.Parameter.Value
}

func logUnableToReadFromSsmError(err error, ssmPath string){
  if err != nil {
    log.Printf("Unable to read from SSM path { %s }: { %v }", ssmPath, err)
  }
}