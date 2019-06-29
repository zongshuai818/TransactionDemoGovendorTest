package Cachelib

import "fmt"

type Cache struct {
	Hash map[string]string
}

func (c *Cache) Set(key string, value string) bool {
	_, ok := c.Hash[key]
	if ok {
		fmt.Println("Keyword exists and its value has been modified ")
	}
	c.Hash[key] = value
	return true
}

func (c *Cache) Get(key string) (string, bool) {
	value, ok := c.Hash[key]
	if !ok {
		value = ""
	}
	return value, ok
}

func (c *Cache) Del(key string) bool {
	_, ok := c.Hash[key]
	if ok {
		delete(c.Hash, key)
	}
	return ok
}
