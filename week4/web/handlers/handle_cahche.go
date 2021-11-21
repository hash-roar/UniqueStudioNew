package handlers

import (
	dbhandlers "authmanager/dbhandle"
	"log"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

func init() {

}

func AddClientAuthToken(clientCode string, userInfo map[string]string) {
	conn := dbhandlers.GetRedisConn()
	vars := redis.Args{}.Add(clientCode).AddFlat(userInfo)
	result, err := conn.Do("HMSET", vars...)
	log.Println(result, err)
}

func GetClientAuthToken(key string, field ...string) ([]string, error) {
	conn := dbhandlers.GetRedisConn()
	vars := redis.Args{}.Add(key).AddFlat(field)
	result, err := redis.Strings(conn.Do("HMGET", vars...))
	return result, err
}

func AddUserToken(uid int, token string) {
	conn := dbhandlers.GetRedisConn()
	conn.Do("SET", "u_token:"+strconv.Itoa(uid), token)
}
func GetUserToken(uid int) (interface{}, error) {
	conn := dbhandlers.GetRedisConn()
	result, err := conn.Do("GET", "u_token:"+strconv.Itoa(uid))
	return result, err
}
