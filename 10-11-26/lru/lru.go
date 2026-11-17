package lru

import "Nov_10_26/list"

type Cache[Key comparable, Value any] struct {
	data map[Key]*list.Node[Value]
	lruChain *list.List[Value]
}

func NewCache[Key comparable, Value any](capacity int) *Cache[Key, Value] {
	return &Cache[Key, Value]{
		data: make(map[Key]*list.Node[Value], capacity),
		lruChain: list.NewList[Value](),
	}
}

func (cache *Cache[Key, Value]) Get(key Key) (Value, bool) {
	node, ok := cache.get(key)
	if !ok {
		var val Value
		return val, false
	}
	return node.Element(), true
}

func (cache *Cache[Key, Value]) get(key Key) (*list.Node[], bool) {
	node, ok := cache.data[key]
	if !ok {
		return nil, false
	}
	cache.lruChain.DetachNode(node)
	cache.lruChain.InsertNodeAsFirst(node)
	return node, true
}

func (cache *Cache[Key, Value]) Set(key Key, val Value) {
	if len(cache.data) == cap(cache.)
}