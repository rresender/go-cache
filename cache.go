package main

// Key generic
type Key interface{}

// Value generic
type Value interface{}

// Entry generic
type Entry struct {
	k Key
	v Value
}

// GeekCache interface
type GeekCache interface {
	Put(e Entry)
	Get(e Entry) Entry
	Remove(e Entry)
	Clear()
}
