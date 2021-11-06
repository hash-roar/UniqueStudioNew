package handler

import (
	"pastebin/cache"
	"pastebin/database/model"
)

func WriteCache(data *model.Pastecode) bool {
	cache.LruCache.Add(data.UrlIndex, data)
	return true
}

func GetCache(key string) (cache.Value, bool) {
	if temp, ok := cache.LruCache.Get(key); ok {
		return temp, true
	}
	return nil, false
}
