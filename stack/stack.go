package stack

import (
	"errors"
	"sync"
)

type stack[T any] struct {
	lock   sync.Mutex
	values []T
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{
		lock:   sync.Mutex{},
		values: make([]T, 0),
	}
}

func (s *stack[T]) secureLock() {
	s.lock.Lock()
	defer s.lock.Unlock()
}

func (s *stack[T]) Push(value T) {
	s.secureLock()
	s.values = append(s.values, value)
}

func (s *stack[T]) Pop() (T, error) {
	s.secureLock()
	lastIndex := len(s.values)
	if lastIndex == 0 {
		var zeroValue T // A variable with the zero value of T
		return zeroValue, errors.New("empty stack")
	}
	value := s.values[lastIndex-1]
	s.values = s.values[:lastIndex-1]
	return value, nil
}

func (s *stack[T]) Peek() T {
	s.secureLock()
	lastIndex := len(s.values)
	value := s.values[lastIndex-1]
	return value
}

func (s *stack[T]) IsEmpty() bool {
	s.secureLock()
	return len(s.values) == 0
}
