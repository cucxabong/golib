package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLength(t *testing.T) {
	q := Queue{}
	assert.Equal(t, 0, q.Length())

	q.EnQueue(10)
	q.EnQueue(20)
	assert.Equal(t, 2, q.Length())
	q.DeQueue()
	assert.Equal(t, 1, q.Length())
	q.Front()
	assert.Equal(t, 1, q.Length())
	q.DeQueue()
	assert.Equal(t, 0, q.Length())
	q.DeQueue()
	assert.Equal(t, 0, q.Length())
}

func TestEnQueue(t *testing.T) {
	q := Queue{}
	q.EnQueue(10)
	assert.Equal(t, 1, q.Length())

	q.EnQueue(20)
	q.EnQueue(30)
	assert.Equal(t, 3, q.Length())
}

func TestDeQueue(t *testing.T) {
	q := Queue{}
	data, err := q.DeQueue()
	assert.Equal(t, ErrorQueueEmpty, err)
	assert.Nil(t, data)

	q.EnQueue(10)
	q.EnQueue(20)
	data, err = q.DeQueue()
	assert.Equal(t, 10, data)
	assert.Nil(t, err)
	assert.Equal(t, 1, q.Length())
}

func TestFront(t *testing.T) {
	q := Queue{}
	data, err := q.Front()
	assert.Equal(t, ErrorQueueEmpty, err)
	assert.Nil(t, data)
	q.EnQueue(10)
	q.EnQueue(20)
	q.EnQueue(30)
	data, err = q.Front()
	assert.Nil(t, err)
	assert.Equal(t, 10, data)
	assert.Equal(t, 3, q.Length())
}
