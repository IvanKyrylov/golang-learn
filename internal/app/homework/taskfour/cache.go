package taskfour

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool // Добавить значение в кэш по ключу.
	Get(key Key) (interface{}, bool)     // Получить значение из кэша по ключу.
	Clear()                              // Очистить кэш.
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	cItem := cacheItem{
		key:   string(key),
		value: value,
	}

	if v, ok := c.items[key]; ok {
		v.Value = cItem
		c.items[key] = v
		c.queue.MoveToFront(v)

		return true
	}

	c.items[key] = c.queue.PushFront(cItem)

	if c.queue.Len() > c.capacity {
		backItem := c.queue.Back()
		c.queue.Remove(backItem)
		delete(c.items, Key(backItem.Value.(cacheItem).key))
	}

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if _, ok := c.items[key]; !ok {
		return nil, false
	}

	getItem := c.items[key]
	c.queue.MoveToFront(getItem)

	return getItem.Value.(cacheItem).value, true
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

type cacheItem struct {
	key   string
	value interface{}
}

func NewCache(capacity int) Cache {
	if capacity < 0 {
		panic("capacity value must be >= 0")
	}

	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
