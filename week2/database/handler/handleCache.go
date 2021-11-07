package handler

import (
	"fmt"
	"pastebin/cache"
	"pastebin/database/model"
)

func WriteCache(data *model.Pastecode) bool {
	cache.LruCache.Add(data.UrlIndex, data)
	return true
}

func GetCache(key string) (cache.Value, bool) {
	fmt.Println(*cache.LruCache)
	if temp, ok := cache.LruCache.Get(key); ok {
		return temp, true
	}
	return nil, false
}
