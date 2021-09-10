package queue

import (
	"fmt"
	"sync"
)

var ErrorQueueEmpty = fmt.Errorf("queue empty")

type Queue struct {
	lock sync.RWMutex
	data []interface{}
}

func (q *Queue) EnQueue(val interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.data = append(q.data, val)
}

func (q *Queue) DeQueue() (interface{}, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.data) == 0 {
		return nil, ErrorQueueEmpty
	}

	data := q.data[0]
	q.data = q.data[1:]

	return data, nil
}

func (q *Queue) Front() (interface{}, error) {
	q.lock.RLock()
	defer q.lock.RUnlock()
	if len(q.data) == 0 {
		return nil, ErrorQueueEmpty
	}

	return q.data[0], nil
}

func (q *Queue) Length() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return len(q.data)
}
