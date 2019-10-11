package main

import (
	"errors"
	"fmt"
)

// LFUHashMap map type
type LFUHashMap map[interface{}]*Node

// LFUCache struct
type LFUCache struct {
	frequencyList   []*List
	cache           LFUHashMap
	lowestFrequency int
	maxFrequency    int
	maxCacheSize    int
	evictionFactor  float32
}

func (c *LFUCache) initFrequencyList() {
	c.frequencyList = make([]*List, c.maxCacheSize)
	for i := 0; i <= c.maxFrequency; i++ {
		c.frequencyList[i] = &List{size: 0}
	}
}

func (c *LFUCache) findNextLowestFrequency() {
	for c.lowestFrequency <= c.maxFrequency && c.frequencyList[c.lowestFrequency].size == 0 {
		c.lowestFrequency++
	}
	if c.lowestFrequency > c.maxFrequency {
		c.lowestFrequency = 0
	}
}

var errLowestFrequency = errors.New("lowest frequency constraint violated")

func (c *LFUCache) doEviction() (err error) {
	currentlyDeleted := 0
	target := int(float32(c.maxCacheSize) * c.evictionFactor)
	for currentlyDeleted < target {
		nodes := c.frequencyList[c.lowestFrequency]
		if nodes.size == 0 {
			err = errLowestFrequency
			return err
		}
		for nodes.size != 0 && currentlyDeleted < target {
			last := nodes.RemoveFromBack()
			delete(c.cache, last)
			currentlyDeleted++
		}
		if nodes.size == 0 {
			c.findNextLowestFrequency()
		}
	}
	return err
}

func (c *LFUCache) put(k, v interface{}) (oldvalue interface{}) {
	oldvalue = nil
	currentNode := c.cache[k]
	if currentNode == nil {
		if len(c.cache) == c.maxCacheSize {
			c.doEviction()
		}
		nodes := c.frequencyList[0]
		nodes.InsertAtFront(v)
		c.cache[k] = nodes.Find(v)
		c.lowestFrequency = 0
	} else {
		oldvalue = currentNode.value
		currentNode.value = v
	}
	return oldvalue
}

func (c *LFUCache) get(k interface{}) interface{} {
	currentNode := c.cache[k]
	if currentNode != nil {
		currentFrequency := currentNode.frequency
		if currentFrequency < c.maxFrequency {
			nextFrequency := currentFrequency + 1
			currentNodes := c.frequencyList[currentFrequency]
			newNodes := c.frequencyList[nextFrequency]
			c.moveToNextFrequency(currentNode, nextFrequency, currentNodes, newNodes)
			c.cache[k] = currentNode
			if c.lowestFrequency == currentFrequency && currentNodes.size == 0 {
				c.lowestFrequency = nextFrequency
			}
		} else {
			nodes := c.frequencyList[currentFrequency]
			nodes.RemoveNode(currentNode)
			nodes.InsertAtFront(currentNode.value)
		}
		return currentNode.value
	}
	return nil
}

func (c *LFUCache) moveToNextFrequency(currentNode *Node, nextFrequency int, currentNodes, newNodes *List) {
	currentNodes.RemoveNode(currentNode)
	newNodes.InsertAtFront(currentNode.value)
	currentNode.frequency = nextFrequency
}

func (c *LFUCache) frequencyOf(k interface{}) int {
	node := c.cache[k]
	if node != nil {
		return node.frequency + 1
	}
	return 0
}

func (c *LFUCache) remove(k interface{}) interface{} {
	currentNode := c.cache[k]
	if currentNode != nil {
		nodes := c.frequencyList[currentNode.frequency]
		nodes.Remove(currentNode)
		if c.lowestFrequency == currentNode.frequency {
			c.findNextLowestFrequency()
		}
		return currentNode.value
	}
	return nil
}

func (c *LFUCache) clear() {
	c.initFrequencyList()
	c.cache = LFUHashMap{}
	c.lowestFrequency = 0
}

func newLFUCache(maxCacheSize int, evictionFactor float32) LFUCache {
	cache := LFUCache{frequencyList: []*List{}, cache: LFUHashMap{}, maxFrequency: maxCacheSize - 1, maxCacheSize: maxCacheSize, lowestFrequency: 0, evictionFactor: evictionFactor}
	cache.initFrequencyList()
	return cache
}

// Put an entry into the cache
func (c *LFUCache) Put(e Entry) {
	c.put(e.k, e.v)
}

// Get an entry from the cache
func (c *LFUCache) Get(e Entry) Entry {
	return Entry{k: e.k, v: c.get(e.k)}
}

// Remove an entry from the cache
func (c *LFUCache) Remove(e Entry) {
	c.remove(e.k)
}

// Clear all the cache
func (c *LFUCache) Clear() {
	c.clear()
}

func main() {
	myLFU := newLFUCache(5, 1)
	myLFU.put(1, 1) // state of myLFU.freq: {1: 1}
	fmt.Printf("F: %d\n", myLFU.frequencyOf(1))
	myLFU.put(2, 2) // state of myLFU.freq: {1: 2<->1}
	fmt.Printf("F: %d\n", myLFU.frequencyOf(2))
	myLFU.put(3, 3) // state of myLFU.freq: {1: 3<->2<->1}
	fmt.Printf("F: %d\n", myLFU.frequencyOf(3))
	myLFU.put(4, 4) // state of myLFU.freq: {1: 4<->3<->2<->1}
	fmt.Printf("F: %d\n", myLFU.frequencyOf(4))
	myLFU.put(5, 5) // state of myLFU.freq: {1: 5<->4<->3<->2<->1}
	fmt.Printf("F: %d\n", myLFU.frequencyOf(5))
	v := myLFU.get(1) // returns 1, state of myLFU.freq: {1: 5<->4<->3<->2, 2:
	println(v.(int))
	fmt.Printf("F: %d\n", myLFU.frequencyOf(1))
	v = myLFU.get(1)
	println(v.(int))
	fmt.Printf("F: %d\n", myLFU.frequencyOf(1))
	v = myLFU.get(1)
	println(v.(int))
	fmt.Printf("F: %d\n", myLFU.frequencyOf(1))
	myLFU.put(6, 6) // state of myLFU.freq: {1: 6<->5<->4<->3, 4: 1}
	v = myLFU.get(6)
	myLFU.put(7, 7) // state of myLFU.freq: {1: 6<->5<->4<->3, 4: 1}
	v = myLFU.get(7)
	println(v.(int))
}
