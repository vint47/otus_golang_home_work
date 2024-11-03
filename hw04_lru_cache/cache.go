package hw04lrucache

import "sync"

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
	keys     map[*ListItem]Key
	mu       sync.RWMutex
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
		keys:     make(map[*ListItem]Key, capacity),
		mu:       sync.RWMutex{},
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	listItem, exists := c.items[key]
	if exists {
		listItem.Value = value
		c.queue.MoveToFront(listItem)

		return true
	}

	if c.queue.Len() == c.capacity {
		lastListItem := c.queue.Back()
		delete(c.items, c.keys[lastListItem])
		delete(c.keys, lastListItem)
		c.queue.Remove(lastListItem)
	}

	listItem = c.queue.PushFront(value)
	c.items[key] = listItem
	c.keys[listItem] = key

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	listItem, exists := c.items[key]
	if !exists {
		return nil, false
	}
	c.queue.MoveToFront(listItem)

	return listItem.Value, exists
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}
