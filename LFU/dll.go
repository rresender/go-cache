package main

import (
	"errors"
	"fmt"
)

// Node obj
type Node struct {
	Value      interface{}
	next, prev *Node
}

// List obj
type List struct {
	size       int
	head, tail *Node
}

var errOutOfBound = errors.New("ERROR - Index out of bound")

// InsertAtFront of List
func (l *List) InsertAtFront(v interface{}) {
	if l.head == nil {
		l.head = &Node{Value: v}
	} else {
		if l.tail == nil {
			l.tail = &Node{Value: v}
			l.head.next = l.tail
			l.tail.prev = l.head
		} else {
			prevHead := l.head
			newHead := &Node{Value: v}
			newHead.next = prevHead
			prevHead.prev = newHead
			l.head = newHead
		}
	}
	l.size++
}

// InsertAtBack of the list
func (l *List) InsertAtBack(v interface{}) {
	if l.head == nil {
		l.head = &Node{Value: v}
		l.tail = nil
	} else {
		if l.tail == nil {
			l.tail = &Node{Value: v}
			l.head.next = l.tail
			l.tail.prev = l.head
		} else {
			prevTail := l.tail
			newTail := &Node{Value: v}
			newTail.prev = prevTail
			prevTail.next = newTail
			l.tail = newTail
		}
	}
	l.size++
}

// InsertNode in the middle of the list
func (l *List) InsertNode(v interface{}, currentNode, nextNode *Node) {
	newNode := &Node{Value: v}
	if nextNode == nil {
		l.tail = newNode
		l.tail.prev = currentNode
		currentNode.next = newNode
	} else {
		currentNode.next = newNode
		nextNode.prev = newNode
		newNode.prev = currentNode
		newNode.next = nextNode
	}
	l.size++
}

// InsertAfter the provide index
func (l *List) InsertAfter(v interface{}, index int) (err error) {
	if index > (l.size - 1) {
		err = errOutOfBound
		return err
	}
	currentNode := l.head
	counter := 0
	for currentNode != nil && counter != index {
		currentNode = currentNode.next
		counter++
	}
	l.InsertNode(v, currentNode, currentNode.next)
	return err
}

// RemoveFromFront of the list
func (l *List) RemoveFromFront() {
	if l.head != nil {
		nextNode := l.head.next
		if nextNode != nil {
			nextNode.prev = nil
		}
		l.head = nil
		l.head = nextNode
		l.size--
	}
}

// RemoveFromBack of the list
func (l *List) RemoveFromBack() (v interface{}) {
	if l.head != nil {
		if l.tail != nil {
			newTail := l.tail.prev
			v = l.tail.Value
			l.tail = nil
			if newTail != nil {
				newTail.next = nil
			}
			l.tail = newTail
		} else {
			v = l.head.Value
			l.head = nil
		}
		l.size--
	}
	return v
}

//RemoveNode in the middle of the list
func (l *List) RemoveNode(nodeToRemove *Node) {
	prevNode := nodeToRemove.prev
	nextNode := nodeToRemove.next
	if prevNode == nil {
		l.head = nil
		l.head = nextNode
		if l.head != nil {
			l.head.prev = nil
		}
	} else if nextNode == nil {
		l.tail = nil
		l.tail = prevNode
		if l.tail.next != nil {
			l.tail.next = nil
		}
	} else {
		nodeToRemove = nil
		prevNode.next = nextNode
		nextNode.prev = prevNode
	}
	l.size--
}

// Remove based on the value
func (l *List) Remove(v interface{}) {
	currentNode := l.head
	for currentNode != nil {
		currentValue := currentNode.Value
		if currentValue == v {
			l.RemoveNode(currentNode)
		}
		currentNode = currentNode.next
	}
}

// Find a Node based on a value
func (l *List) Find(v interface{}) *Node {
	currentNode := l.head
	for currentNode != nil {
		currentValue := currentNode.Value
		if currentValue == v {
			return currentNode
		}
		currentNode = currentNode.next
	}
	return nil
}

// Print the list
func (l *List) Print() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Printf("%v ", currentNode.Value)
		currentNode = currentNode.next
	}
}
