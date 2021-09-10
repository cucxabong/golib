package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	s := Stack{}
	assert.Equal(t, 0, s.Len())

	s.Push(10)
	assert.Equal(t, 1, s.Len())

	s.Push(20)
	assert.Equal(t, 2, s.Len())

	s.Top()
	assert.Equal(t, 2, s.Len())

	s.Pop()
	assert.Equal(t, 1, s.Len())

	_, err := s.Pop()
	assert.Equal(t, 0, s.Len())
	assert.Nil(t, err)

	_, err = s.Pop()
	assert.NotNil(t, err)
}

func TestPush(t *testing.T) {
	s := Stack{}
	s.Push(123)

	data, err := s.Pop()
	assert.Equal(t, 123, data)
	assert.Nil(t, err)

	_, err = s.Pop()
	assert.Equal(t, err, ErrorEmptyStack)
}

func TestPop(t *testing.T) {
	s := Stack{}
	_, err := s.Pop()
	assert.Equal(t, ErrorEmptyStack, err)

	s.Push(123)
	s.Push(456)
	data, err := s.Pop()
	assert.Equal(t, 456, data)
	assert.Equal(t, 1, s.Len())
}

func TestTop(t *testing.T) {
	s := Stack{}
	_, err := s.Top()
	assert.Equal(t, ErrorEmptyStack, err)

	s.Push(10)
	s.Push(20)
	data, err := s.Top()
	assert.Equal(t, 20, data)
	assert.Nil(t, err)
	assert.Equal(t, 2, s.Len())
}
