package handler

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"pastebin/database/model"

	"github.com/garyburd/redigo/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var pool *redis.Pool

var postgresqlDb *gorm.DB
var dsn string = "user=dbuser password=zlf dbname=pastebin port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func init() {
	pool = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	// dsn := "user=dbuser password=zlf dbname=pastebin port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// postgresqlDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// sqlDb, err2 := postgresqlDb.DB()
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	// sqlDb.SetMaxIdleConns(10)
	// sqlDb.SetMaxOpenConns(100)
	// sqlDb.SetConnMaxLifetime(time.Second * 300)
	// println("db init success")
}

func WriteRedis(data *model.Pastecode) (md5str string) {
	c := pool.Get()
	defer c.Close()
	data_json, _ := json.Marshal(data)
	data_md5 := md5.Sum(data_json)
	c.Do("Set", hex.EncodeToString(data_md5[:]), data_json)
	md5str = hex.EncodeToString(data_md5[:])
	return
}

func ReadRedis(paste_index string) (content *model.Pastecode, err error) {
	c := pool.Get()
	defer c.Close()
	temp, err := redis.Bytes(c.Do("Get", paste_index))
	if err != nil {
		fmt.Println(err)
	}
	content = &model.Pastecode{}
	json.Unmarshal(temp, content)
	return
}

func Writepostgresql(data *model.Pastecode) {

}

func Readpostgresql(url_index string) *model.Pastecode {
	// sqlDB, err := postgresqlDb.DB()
	// if err != nil {
	// 	fmt.Println("connect db server failed.")
	// }
	// if err := sqlDB.Ping(); err != nil {
	// 	sqlDB.Close()
	// }
	// data := model.Pastecode{}
	// // db,err := postgresqlDb.DB()
	// // if err!=nil {
	// // 	fmt.Println(err)
	// // }
	// result := postgresqlDb.First(&data)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }
	// fmt.Println(data)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	data := model.Pastecode{}
	db.AutoMigrate(&data)
	db.Create(&model.Pastecode{Content: "content", Syntax: "plaintext", Poster: "zlf", Expiration: "n"})
	var readdata model.Pastecode
	db.First(&readdata)
	fmt.Println(readdata)
	return &readdata

}
