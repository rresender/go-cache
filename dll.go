package main

import (
	"errors"
	"fmt"
)

// Node obj
type Node struct {
	frequency  int
	value      interface{}
	next, prev *Node
}

// List obj
type List struct {
	size       int
	head, tail *Node
}

var errOutOfBound = errors.New("ERROR - Index out of bound")

// InsertAtFront of List
func (l *List) InsertAtFront(v interface{}) (n *Node) {
	n = &Node{value: v}
	if l.head == nil {
		l.head = n
	} else {
		if l.tail == nil {
			l.tail = n
			l.head.next = l.tail
			l.tail.prev = l.head
		} else {
			prevHead := l.head
			newHead := n
			newHead.next = prevHead
			prevHead.prev = newHead
			l.head = newHead
		}
	}
	l.size++
	return n
}

// InsertAtBack of the list
func (l *List) InsertAtBack(v interface{}) (n *Node) {
	n = &Node{value: v}
	if l.head == nil {
		l.head = n
		l.tail = nil
	} else {
		if l.tail == nil {
			l.tail = n
			l.head.next = l.tail
			l.tail.prev = l.head
		} else {
			prevTail := l.tail
			newTail := n
			newTail.prev = prevTail
			prevTail.next = newTail
			l.tail = newTail
		}
	}
	l.size++
	return n
}

// InsertNode in the middle of the list
func (l *List) InsertNode(v interface{}, currentNode, nextNode *Node) (n *Node) {
	n = &Node{value: v}
	if nextNode == nil {
		l.tail = n
		l.tail.prev = currentNode
		currentNode.next = n
	} else {
		currentNode.next = n
		nextNode.prev = n
		n.prev = currentNode
		n.next = nextNode
	}
	l.size++
	return n
}

// InsertAfter the provide index
func (l *List) InsertAfter(v interface{}, index int) (n *Node, err error) {
	if index > (l.size - 1) {
		err = errOutOfBound
		return nil, err
	}
	currentNode := l.head
	counter := 0
	for currentNode != nil && counter != index {
		currentNode = currentNode.next
		counter++
	}
	n = l.InsertNode(v, currentNode, currentNode.next)
	return n, err
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
			v = l.tail.value
			l.tail = nil
			if newTail != nil {
				newTail.next = nil
			}
			l.tail = newTail
		} else {
			v = l.head.value
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
func (l *List) Remove(v interface{}) (n *Node) {
	currentNode := l.head
	for currentNode != nil {
		currentvalue := currentNode.value
		if currentvalue == v {
			n = currentNode
			l.RemoveNode(currentNode)
		}
		currentNode = currentNode.next
	}
	return n
}

// Find a Node based on a value
func (l *List) Find(v interface{}) *Node {
	currentNode := l.head
	for currentNode != nil {
		currentvalue := currentNode.value
		if currentvalue == v {
			return currentNode
		}
		currentNode = currentNode.next
	}
	return nil
}

// Print the list
func (l *List) Print() {
	currentNode := l.head
	fmt.Print("[")
	for currentNode != nil {
		fmt.Printf("%v", currentNode.value)
		currentNode = currentNode.next
		if currentNode != nil {
			fmt.Print("->")
		}
	}
	fmt.Print("]\n")
}

func (l *List) isEmpty() bool {
	return l.size == 0
}
