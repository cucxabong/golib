package stack

import (
	"fmt"
	"sync"
)

var ErrorEmptyStack = fmt.Errorf("empty stack")

type Stack struct {
	data []interface{}
	lock sync.RWMutex
}

func (s *Stack) Len() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.data)
}

func (s *Stack) Push(item interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = append(s.data, item)
}

func (s *Stack) Pop() (interface{}, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.data) == 0 {
		return nil, ErrorEmptyStack
	}

	data := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return data, nil
}

func (s *Stack) Top() (interface{}, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if len(s.data) == 0 {
		return nil, ErrorEmptyStack
	}

	return s.data[len(s.data)-1], nil
}
