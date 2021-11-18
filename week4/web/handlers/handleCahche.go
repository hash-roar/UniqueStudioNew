package handlers

import dbhandlers "authmanager/dbhandle"

func init() {

}

func AddClientAuthToken(clientCode string, userInfo string) {
	conn := dbhandlers.GetRedisConn()
	conn.Do("SET", clientCode, userInfo)
}

func GetClientAuthToken(key string) (interface{}, error) {
	conn := dbhandlers.GetRedisConn()
	result, err := conn.Do("GET", key)
	return result, err
}
