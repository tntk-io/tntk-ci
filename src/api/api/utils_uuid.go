package api

import "github.com/satori/uuid"

func getUuid() string{
  return uuid.NewV4().String()
}
