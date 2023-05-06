package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection(dsn string) (*gorm.DB, error) {
	//dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
