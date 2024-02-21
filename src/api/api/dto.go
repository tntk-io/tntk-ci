package api

import (
  "fmt"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
  "log"
)

var gormDb *gorm.DB

func getGormDb() *gorm.DB{
  if gormDb == nil {
    gormDb, err = gorm.Open(mysql.Open(getDatabaseConnectionInfo()), &gorm.Config{})
    if err != nil {
      log.Panicf("failed to connect to database: { %s }", getDatabaseConnectionInfo())
    }
    err := gormDb.AutoMigrate(&User{})
    if err != nil {
      log.Printf("Error automatically creating User table: %v", err)
    }
  }

  return gormDb
}

func getDatabaseConnectionInfo() string {
  //return fmt.Sprintf("host=%s port=%s user=%s "+
  //  "password=%s dbname=%s sslmode=disable",
  //  getDbHost(), getDbPort(), getDbUsername(), getDbPassword(), getDbName())
  return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=5s",getDbUsername(), getDbPassword(), getDbHost(), getDbPort(), getDbName())
}