package dbhandlers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var dsn string = "user=dbuser password=zlf dbname=pastebin port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func init() {
	db, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func GetDbConn() *gorm.DB {
	return db
}
