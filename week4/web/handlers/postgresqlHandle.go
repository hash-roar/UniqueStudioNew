package handlers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var dsn string = "user=dbuser password=zlf dbname=pastebin port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func init() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("open database error")
	}
}

func GetDbConn() *gorm.DB {
	return db
}
