package cache

import "container/list"

type Cache struct {
	maxBytes  int64
	usedBytes int64
	ll        *list.List
	cache     map[string]*list.Element
	del       *list.List
	callback  func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int64
}

var LruCache *Cache

// func init() {
// 	LruCache = initCache(1024 * 1024 * 20,func(key string, value Value) {})
// }

func InitCache(maxBytes int64, callback func(key string, value Value)) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
		callback: callback,
	}
}

func (ca *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := ca.cache[key]; ok {
		ca.ll.MoveToFront(ele)
		result := ele.Value.(*entry)
		return result.value, true
	}
	return
}

func (ca *Cache) Delete(key string) bool {
	if ele, ok := ca.cache[key]; ok {
		temp := ele.Value.(*entry)
		ca.ll.Remove(ele)
		delete(ca.cache, temp.key)
		ca.usedBytes -= int64(len(temp.key)) + int64(temp.value.Len())
		if ca.callback != nil {
			ca.callback(temp.key, temp.value)
		}
		return true
	}
	return false
}

func (ca *Cache) DeleteBack() {
	if ele := ca.ll.Back(); ele != nil {
		result := ele.Value.(*entry)
		ca.ll.Remove(ele)
		delete(ca.cache, result.key)
		ca.usedBytes -= int64(len(result.key)) + int64(result.value.Len())
		if ca.callback != nil {
			ca.callback(result.key, result.value)
		}
	}

}

func (ca *Cache) Add(key string, value Value) {
	if ele, ok := ca.cache[key]; ok {
		ca.ll.MoveToFront(ele)
		result := ele.Value.(*entry)
		ca.usedBytes += int64(value.Len()) - int64(result.value.Len())
	} else {
		ele := ca.ll.PushFront(&entry{key: key, value: value})
		ca.cache[key] = ele
		ca.usedBytes += int64(len(key)) + int64(value.Len())
	}
	for ca.maxBytes > 0 && ca.maxBytes < ca.usedBytes {
		ca.DeleteBack()
	}

}
