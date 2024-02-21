package api

import "os"

func getDbHost() string{
  return readValueFromSsmIfSsmPathProvided(os.Getenv("DB_HOST"))
}

func getDbPort() string{
  return readValueFromSsmIfSsmPathProvided(os.Getenv("DB_PORT"))
}

func getDbUsername() string{
  return readValueFromSsmIfSsmPathProvided(os.Getenv("DB_USERNAME"))
}

func getDbPassword() string{
  return readValueFromSsmIfSsmPathProvided(os.Getenv("DB_PASSWORD"))
}

func getDbName() string{
  return readValueFromSsmIfSsmPathProvided(os.Getenv("DB_NAME"))
}