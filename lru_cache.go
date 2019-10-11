package main

import "fmt"

// LRUHashMap map type
type LRUHashMap map[interface{}]interface{}

// LRUCache struct
type LRUCache struct {
	keys         List
	cache        LRUHashMap
	maxCacheSize int
}

func newLRUCache(n int) LRUCache {
	return LRUCache{keys: List{size: 0}, cache: LRUHashMap{}, maxCacheSize: n}
}

func (c *LRUCache) put(k interface{}, v interface{}) {
	if _, found := c.cache[k]; found {
		c.keys.Remove(k)
	} else {
		if c.keys.size == c.maxCacheSize {
			if last := c.keys.RemoveFromBack(); last != nil {
				delete(c.cache, last)
			}
		}
	}
	c.keys.InsertAtFront(k)
	c.cache[k] = v
}

func (c *LRUCache) get(k interface{}) interface{} {
	if _, found := c.cache[k]; !found {
		return nil
	}
	node := c.keys.Find(k)
	if node == nil {
		return nil
	}
	c.keys.RemoveNode(node)
	c.keys.InsertAtFront(k)
	return node.value
}

func (c *LRUCache) clear() {
	n := newLRUCache(c.maxCacheSize)
	c = nil
	c = &n
}

func (c *LRUCache) displayKeys() {
	c.keys.Print()
}

func (c *LRUCache) displayValue() {
	fmt.Print("[")
	first := true
	for k, v := range c.cache {
		if !first {
			fmt.Printf(", ")
		} else {
			first = false
		}
		fmt.Printf("%v->%v", k, v)
	}
	fmt.Print("]\n")
}

func main1() {
	cache := newLRUCache(4)
	cache.put(1, 1)
	cache.put(2, 2)
	cache.put(3, 3)
	cache.put(1, 1)
	cache.put(4, 4)
	cache.put(5, 5)
	cache.put("hello", "World")
	v := cache.get("hello").(string)
	cache.displayKeys()
	cache.displayValue()
	println(v)
}
