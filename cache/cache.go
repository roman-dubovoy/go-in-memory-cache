package cache

type BasicCacheInterface interface {
	Set(key string, value interface{})
	Get(key string)
	Delete(key string)
}

type InMemoryCache struct {
	memory map[string]interface{}
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		memory: make(map[string]interface{}),
	}
}

func (c InMemoryCache) Set(key string, value interface{}) {
	c.memory[key] = value
}

func (c InMemoryCache) Get(key string) interface{} {
	return c.memory[key]
}

func (c InMemoryCache) Delete(key string) {
	delete(c.memory, key)
}
