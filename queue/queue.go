package queue

import (
	"errors"
	"sync"
)

type queue[T any] struct {
	lock   sync.Mutex
	values []T
}

func NewQueue[T any]() *queue[T] {
	return &queue[T]{
		lock:   sync.Mutex{},
		values: make([]T, 0),
	}
}

func (q *queue[T]) secureLock() {
	q.lock.Lock()
	defer q.lock.Unlock()
}

func (q *queue[T]) Enqueue(value T) {
	q.secureLock()
	q.values = append(q.values, value)

}

func (q *queue[T]) Dequeue() (T, error) {
	q.secureLock()
	length := len(q.values)

	if length == 0 {
		var zeroValue T
		return zeroValue, errors.New("empty queue")
	}

	value := q.values[0]
	q.values = q.values[1:]

	return value, nil
}

func (q *queue[T]) Length() int {
	q.secureLock()
	return len(q.values)

}
