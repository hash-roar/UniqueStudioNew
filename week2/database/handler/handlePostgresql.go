package handler

import (
	"fmt"
	"pastebin/database/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn string = "user=dbuser password=zlf dbname=pastebin port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func init() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&model.Pastecode{})
}

func Addpastedata(data *model.Pastecode) (rowAffect int64, err error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	result := db.Create(data)
	return result.RowsAffected, result.Error
}

func Getpastedata(urlIndex string) (*model.Pastecode, int64) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	data := model.Pastecode{}
	result := db.Where("url_index = ?", urlIndex).First(&data)
	return &data, result.RowsAffected
}
