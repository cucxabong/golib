package linkedlist

import (
	"fmt"
	"sync"
)

type node struct {
	next *node
	data int
}

type LinkedList struct {
	head *node
	lock *sync.RWMutex
}

var (
	ErrorItemNotFound = fmt.Errorf("item not found")
)

func NewLinkedList() *LinkedList {
	return &LinkedList{
		lock: &sync.RWMutex{},
	}
}

// String return a formatted string contain all node in the linked list
func (ll LinkedList) String() string {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	s := ""
	iter := ll.head
	for iter != nil {
		s += fmt.Sprint(iter.data)
		iter = iter.next
		if iter != nil {
			s += " -> "
		}
	}
	return s
}

// Append insert a new item at the end of a linked list
func (ll *LinkedList) Append(val int) {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	if ll.head == nil {
		ll.head = &node{
			data: val,
		}
	} else {
		iter := ll.head
		for iter.next != nil {
			iter = iter.next
		}

		iter.next = &node{
			data: val,
		}
	}
}

// Insert insert a new item at the beginning of a linked list
func (ll *LinkedList) Insert(val int) {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	newNode := &node{
		data: val,
	}

	if ll.head == nil {
		ll.head = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}
}

// Remove delete the first occurance of val from a linked list
func (ll *LinkedList) Remove(val int) error {
	ll.lock.Lock()
	defer ll.lock.Unlock()

	if ll.head == nil {
		return ErrorItemNotFound
	}

	if ll.head.data == val && ll.head.next == nil {
		ll.head = nil
		return nil
	}

	iter := ll.head
	var prev *node

	for iter != nil {
		if iter.data == val {
			if prev == nil {
				ll.head = iter.next
				return nil
			}

			prev.next = iter.next
			return nil
		}

		prev = iter
		iter = iter.next
	}

	return ErrorItemNotFound
}

// Length return the number of items in a linked list
func (ll *LinkedList) Length() int {
	ll.lock.RLock()
	defer ll.lock.RUnlock()
	c := 0
	iter := ll.head

	for iter != nil {
		c++
		iter = iter.next
	}

	return c
}

// Reverse reverse the given linked list
func (ll *LinkedList) Reverse() {
	var prev *node
	ll.lock.Lock()
	defer ll.lock.Unlock()

	if ll != nil {
		for ll.head != nil {
			next := ll.head.next
			ll.head.next = prev
			prev = ll.head
			ll.head = next
		}
	}

	ll.head = prev
}
