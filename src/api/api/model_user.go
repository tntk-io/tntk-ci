package api

import (
  "gorm.io/gorm"
)

type User struct {
  gorm.Model
  ID       uint   `gorm:"primaryKey"`
  Name     string `gorm:"uniqueIndex;size:256"`
  Password string
  Token    string `gorm:"uniqueIndex;size:256"`
}

func (user User) IsPasswordMatch(password string) bool{
  return GetMd5FromString(password) == user.Password
}