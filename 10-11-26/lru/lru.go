package lru

import "Nov_10_26/list"

type pair[Key comparable, Value any] struct {
	key Key
	val Value
}

type Cache[Key comparable, Value any] struct {
	data     map[Key]*list.Node[pair[Key, Value]]
	lruChain *list.List[pair[Key, Value]]
	capacity int
}

func NewCache[Key comparable, Value any](capacity int) *Cache[Key, Value] {
	return &Cache[Key, Value]{
		data:     make(map[Key]*list.Node[pair[Key, Value]], capacity),
		lruChain: list.NewList[pair[Key, Value]](),
		capacity: capacity,
	}
}

func (cache *Cache[Key, Value]) Get(key Key) (Value, bool) {
	node, ok := cache.get(key)
	if !ok {
		var val Value
		return val, false
	}
	return node.Element().val, true
}

func (cache *Cache[Key, Value]) get(key Key) (*list.Node[pair[Key, Value]], bool) {
	node, ok := cache.data[key]
	if !ok {
		return nil, false
	}
	cache.lruChain.DetachNode(node)
	cache.lruChain.InsertNodeAsFirst(node)
	return node, true
}

func (cache *Cache[Key, Value]) Set(key Key, val Value) {
	node, ok := cache.get(key)
	if ok {
		node.SetElement(pair[Key, Value]{key: key, val: val})
		return
	}
	if len(cache.data) == cache.capacity {
		pair, ok := cache.lruChain.RemoveBack()
		if !ok {
			panic("unexpected empty lru chain")
		}
		delete(cache.data, pair.key)
	}
	pair := pair[Key, Value]{key: key, val: val}
	cache.lruChain.AddFront(pair)
	cache.data[key] = cache.lruChain.First()
}
