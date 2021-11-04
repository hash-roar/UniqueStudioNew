package cache

import "container/list"

type Cache struct {
	maxBytes  int64
	usedBytes int64
	ll        *list.List
	cache     map[string]*list.Element
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func initCache(maxBytes int64) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
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

func (ca Cache) DeleteBack(key string) {
	if ele := ca.ll.Back(); ele != nil {
		result := ele.Value.(*entry)
		ca.ll.Remove(ele)
		delete(ca.cache, result.key)
		ca.usedBytes -= int64(len(result.key)) + int64(result.value.Len())
	}

}
