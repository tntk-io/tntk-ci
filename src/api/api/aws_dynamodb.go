package api

import (
  "context"
  "encoding/json"
  "fmt"
  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
  "github.com/aws/aws-sdk-go-v2/service/dynamodb"
  "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
  "github.com/thoas/go-funk"
  "log"
  "os"
  "time"
)

var dynamodbClient *dynamodb.Client

func getDynamodbClient() *dynamodb.Client{
  if dynamodbClient == nil {
    dynamodbClient = dynamodb.NewFromConfig(NewAwsConfig())
  }
  return dynamodbClient
}

func CreateEntryInDynamoDB(username string, fileNames []string) {

  jsonBytes, _ := json.Marshal(fileNames)

  out, err := getDynamodbClient().PutItem(context.TODO(), &dynamodb.PutItemInput{
    TableName: aws.String(getDynamodbTableName()),
    Item: map[string]types.AttributeValue{
      "username": &types.AttributeValueMemberS{Value: username},
      "files":    &types.AttributeValueMemberS{Value: string(jsonBytes)},
      "created": &types.AttributeValueMemberN{Value: fmt.Sprintf("%d", time.Now().Unix())},
    },
  })

  if err != nil {
    panic(err)
  }

  fmt.Println(out.Attributes)

}

func createDynamodbTableIfMissing(tableName string){
  if isDynamodbTableMissing(tableName) {
    cfg := NewAwsConfig()
    dynamodbClient := dynamodb.NewFromConfig(cfg)
    _, err := dynamodbClient.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
      AttributeDefinitions: []types.AttributeDefinition{
        {
          AttributeName: aws.String("username"),
          AttributeType: types.ScalarAttributeTypeS,
        },
        //{
        //  AttributeName: aws.String("files"),
        //  AttributeType: types.ScalarAttributeTypeS,
        //},
      },
      KeySchema: []types.KeySchemaElement{
        {
          AttributeName: aws.String("username"),
          KeyType:       types.KeyTypeHash,
        },
        //{
        //  AttributeName: aws.String("files"),
        //  KeyType:       types.KeyTypeRange,
        //},
      },
      TableName:   aws.String(tableName),
      BillingMode: types.BillingModePayPerRequest,
    })
    if err != nil {
      panic(err)
    }
  }

}

func isDynamodbTableMissing(tableName string) bool{
  return ! isDynamodbTableExists(tableName)
}

func isDynamodbTableExists(tableName string) bool{
  cfg := NewAwsConfig()
  dynamodbClient := dynamodb.NewFromConfig(cfg)
  dynamodbTablesListOutput, err := dynamodbClient.ListTables(context.TODO(), &dynamodb.ListTablesInput{
    //ExclusiveStartTableName: aws.String(tableName),
  })
  if err != nil {
    log.Printf("Unable to list Dynamodb tables: %v", err)
    return false
  }

  for _, tableNameFromOutput := range dynamodbTablesListOutput.TableNames {
    if tableNameFromOutput == tableName{
      return true
    }
  }

  return false
}

func getEntryFromDynamoDBByUsername(username string) map[string]types.AttributeValue{

  cfg := NewAwsConfig()

  dynamodbClient := dynamodb.NewFromConfig(cfg)
  out, err := dynamodbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
    TableName: aws.String(getDynamodbTableName()),
    Key: map[string]types.AttributeValue{
      "username": &types.AttributeValueMemberS{Value: username},
    },
  })

  if err != nil {
    panic(err)
  }

  return out.Item
}

func getDynamodbTableName()string {
  return readValueFromSsmIfSsmPathProvided(os.Getenv("DYNAMODB_TABLE_NAME"))
}

func updateEntryInDynamoDb(username, filenamesJsonString string, statusCode int){

  if statusCode >= 200 {
    var filenamesInDynamoDbJsonString string
    var filenamesInDynamoDb []string

    var filenames map[string]string

    json.Unmarshal([]byte(filenamesJsonString), &filenames)

    entry := getEntryFromDynamoDBByUsername(username)
    filesInDynamoDb := entry["files"]
    err := attributevalue.Unmarshal(filesInDynamoDb, &filenamesInDynamoDbJsonString)
    if err != nil {
      fmt.Println(err)
    }
    json.Unmarshal([]byte(filenamesInDynamoDbJsonString), &filenamesInDynamoDb)

    //fmt.Println(fmt.Sprintf("%v", filenamesInDynamoDbJsonString))

    for _, filename := range funk.Values(filenames).([]string) {
      if ! funk.Contains(filenamesInDynamoDbJsonString, filename){
        filenamesInDynamoDb = append(filenamesInDynamoDb, filename)
      }
    }

    CreateEntryInDynamoDB(username, filenamesInDynamoDb)

  }
}

func removeFilenameFromEntryInDynamoDb(username, filename string){

    var filenamesInDynamoDbJsonString string
    var filenamesInDynamoDb []string

    entry := getEntryFromDynamoDBByUsername(username)
    filesInDynamoDb := entry["files"]
    err := attributevalue.Unmarshal(filesInDynamoDb, &filenamesInDynamoDbJsonString)
    if err != nil {
      fmt.Println(err)
    }
    json.Unmarshal([]byte(filenamesInDynamoDbJsonString), &filenamesInDynamoDb)

    CreateEntryInDynamoDB(username, funk.Subtract(filenamesInDynamoDb, []string{filename}).([]string))

}

func GetFilesListAttributeFromDynamoDbEntryByUsername(username string) (jsonString string){

  entry := getEntryFromDynamoDBByUsername(username)
  filesInDynamoDb := entry["files"]
  err := attributevalue.Unmarshal(filesInDynamoDb, &jsonString)
  if err != nil {
    fmt.Println(err)
    return
  }

  return jsonString
}
