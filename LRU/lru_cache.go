package main

// Hash map type
type Hash map[interface{}]bool

// LRUCache struct
type LRUCache struct {
	Queue List
	Hash  Hash
	Size  int
}

func newLRUCache(n int) LRUCache {
	return LRUCache{Queue: List{size: 0}, Hash: Hash{}, Size: n}
}

func (c *LRUCache) refer(key int) {
	if found := c.Hash[key]; found {
		c.Queue.Remove(key)
	} else {
		if c.Queue.size == c.Size {
			if last := c.Queue.RemoveFromBack(); last != nil {
				delete(c.Hash, last)
			}
		}
	}
	c.Queue.InsertAtFront(key)
	c.Hash[key] = true
}

// display contents of cache
func (c *LRUCache) display() {
	c.Queue.Print()
}

func main() {
	cache := newLRUCache(4)
	cache.refer(1)
	cache.refer(2)
	cache.refer(3)
	cache.refer(1)
	cache.refer(4)
	cache.refer(5)
	cache.display()
}
