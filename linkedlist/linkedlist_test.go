package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppend(t *testing.T) {
	ll := NewLinkedList()

	ll.Append(10)
	assert.Equal(t, 1, ll.Length())

	ll.Append(20)
	ll.Append(30)
	assert.Equal(t, 3, ll.Length())
	assert.Equal(t, "10 -> 20 -> 30", ll.String())
}

func TestInsert(t *testing.T) {
	ll := NewLinkedList()
	ll.Insert(10)
	assert.Equal(t, 1, ll.Length())

	ll.Insert(20)
	ll.Insert(30)
	assert.Equal(t, 3, ll.Length())
	assert.Equal(t, "30 -> 20 -> 10", ll.String())
}

func TestRemove(t *testing.T) {
	ll := NewLinkedList()
	err := ll.Remove(10)
	assert.Equal(t, ErrorItemNotFound, err)

	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	err = ll.Remove(20)
	assert.Equal(t, 2, ll.Length())
	assert.Nil(t, err)
	assert.Equal(t, "10 -> 30", ll.String())
}

func TestLength(t *testing.T) {
	ll := NewLinkedList()
	assert.Equal(t, 0, ll.Length())

	ll.Insert(10)
	assert.Equal(t, 1, ll.Length())
	ll.Append(20)
	assert.Equal(t, 2, ll.Length())

	ll.Remove(10)
	assert.Equal(t, 1, ll.Length())
	ll.Remove(100)
	assert.Equal(t, 1, ll.Length())
}

func TestReverse(t *testing.T) {
	ll := NewLinkedList()
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	ll.Reverse()

	assert.Equal(t, 3, ll.Length())
	assert.Equal(t, "30 -> 20 -> 10", ll.String())

}
