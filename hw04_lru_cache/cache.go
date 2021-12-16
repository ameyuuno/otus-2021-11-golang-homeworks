package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	item, exists := c.items[key]
	cachingItem := cacheItem{key, value}

	if exists {
		item.Value = cachingItem
		c.queue.MoveToFront(item)

		return true
	}

	if c.queue.Len() == c.capacity {
		itemToRemove := c.queue.Back()

		c.queue.Remove(itemToRemove)
		delete(c.items, itemToRemove.Value.(cacheItem).key)
	}

	c.items[key] = c.queue.PushFront(cachingItem)

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	item, exists := c.items[key]

	if exists {
		c.queue.MoveToFront(item)
		return item.Value.(cacheItem).value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	*c = lruCache{
		capacity: c.capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, c.capacity),
	}
}
